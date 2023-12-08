package websocket

import (
	"context"
	"flag"
	"fmt"
	"github.com/gobwas/ws/wsutil"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gobwas/httphead"
	"github.com/gobwas/ws"
)

func StartWS(port uint) {
	log.SetFlags(0)
	flag.Parse()

	http.HandleFunc("/ws/", wsHandler)

	addr := ":" + strconv.FormatUint(uint64(port), 10)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen %q error: %v", addr, err)
	}
	log.Printf("WS listening %s", ln.Addr())

	var (
		s     = new(http.Server)
		serve = make(chan error, 1)
		sig   = make(chan os.Signal, 1)
	)
	signal.Notify(sig, syscall.SIGTERM)
	go func() { serve <- s.Serve(ln) }()

	select {
	case err := <-serve:
		log.Fatal(err)
	case sig := <-sig:
		const timeout = 5 * time.Second

		log.Printf("signal %q received; shutting down with %s timeout", sig, timeout)

		ctx, _ := context.WithTimeout(context.Background(), timeout)
		if err := s.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}
}

var (
	closeInvalidPayload = ws.MustCompileFrame(
		ws.NewCloseFrame(ws.NewCloseFrameBody(
			ws.StatusInvalidFramePayloadData, "",
		)),
	)
	closeProtocolError = ws.MustCompileFrame(
		ws.NewCloseFrame(ws.NewCloseFrameBody(
			ws.StatusProtocolError, "",
		)),
	)
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	newConnection := *NewWsConnection()

	u := ws.HTTPUpgrader{
		Extension: func(opt httphead.Option) bool {
			log.Printf("extension: %s", opt)
			return false
		},
	}

	_, joinErr := newConnection.OnAuth(r)

	if joinErr != nil {
		w.Write([]byte("accessToken not found"))
		w.WriteHeader(400)
		return
	}

	conn, _, _, err := u.Upgrade(r, w)

	if err != nil {
		log.Printf("upgrade error: %s", err)
		return
	}
	newWs := &WsRoom{}

	newConnection.AddConnection(&conn)
	AddWsConnection(&newConnection)

	defer conn.Close()

	state := ws.StateServerSide

	textPending := false
	utf8Reader := wsutil.NewUTF8Reader(nil)
	cipherReader := wsutil.NewCipherReader(nil, [4]byte{0, 0, 0, 0})

	for {
		header, err := ws.ReadHeader(conn)

		if err != nil {
			log.Printf("read header error: %s", err)
			break
		}
		if err = ws.CheckHeader(header, state); err != nil {
			log.Printf("header check error: %s", err)
			conn.Write(closeProtocolError)
			return
		}

		cipherReader.Reset(
			io.LimitReader(conn, header.Length),
			header.Mask,
		)

		var utf8Fin bool
		var r io.Reader = cipherReader

		switch header.OpCode {
		case ws.OpPing:
			header.OpCode = ws.OpPong
			header.Masked = false
			ws.WriteHeader(conn, header)
			io.CopyN(conn, cipherReader, header.Length)
			continue

		case ws.OpPong:
			io.CopyN(ioutil.Discard, conn, header.Length)
			continue

		case ws.OpClose:
			newConnection.DisconnectCallback()
			utf8Fin = true

		case ws.OpContinuation:
			if textPending {
				utf8Reader.Source = cipherReader
				r = utf8Reader
			}
			if header.Fin {
				state = state.Clear(ws.StateFragmented)
				textPending = false
				utf8Fin = true
			}

		case ws.OpText:
			utf8Reader.Reset(cipherReader)
			r = utf8Reader

			if !header.Fin {
				state = state.Set(ws.StateFragmented)
				textPending = true
			} else {
				utf8Fin = true
			}

		case ws.OpBinary:
			if !header.Fin {
				state = state.Set(ws.StateFragmented)
			}
		}

		payload := make([]byte, header.Length)
		_, err = io.ReadFull(r, payload)
		if err == nil && utf8Fin && !utf8Reader.Valid() {
			err = wsutil.ErrInvalidUTF8
		}
		if err != nil {
			log.Printf("read payload error: %s", err)
			if err == wsutil.ErrInvalidUTF8 {
				conn.Write(closeInvalidPayload)
			} else {
				conn.Write(ws.CompiledClose)
			}
			return
		}

		if header.OpCode == ws.OpClose {
			code, reason := ws.ParseCloseFrameData(payload)
			log.Printf("close frame received: %v %v", code, reason)

			if !code.Empty() {
				switch {
				case code.IsProtocolSpec() && !code.IsProtocolDefined():
					err = fmt.Errorf("close code from spec range is not defined")
				default:
					err = ws.CheckCloseFrameData(code, reason)
				}
				if err != nil {
					log.Printf("invalid close data: %s", err)
					conn.Write(closeProtocolError)
				} else {
					ws.WriteFrame(conn, ws.NewCloseFrame(ws.NewCloseFrameBody(
						code, "",
					)))
				}
				return
			}

			conn.Write(ws.CompiledClose)
			return
		}

		HandleEvent(&newConnection, payload, newWs)
	}
}
