package docker2dot_test

import (
	"reflect"
	"testing"

	"github.com/po3rin/dockerdot/docker2dot"
)

func TestDocker2Dot(t *testing.T) {
	tests := []struct {
		input []byte
		want  []byte
	}{
		{
			input: []byte(
				`FROM golang:1.12 AS stage0
WORKDIR /go
ADD ./ /go
RUN go build -o stage0_bin
FROM golang:1.12 AS stage1
WORKDIR /go
ADD ./ /go
RUN go build -o stage1_bin
FROM golang:1.12
COPY --from=stage0 /go/stage0_bin /
COPY --from=stage1 /go/stage1_bin / `,
			),
			want: []byte(
				`digraph {
  "sha256:1c5320070ff30eecf3265f227b2646e54428945092b0866a55da4bb20415f066" [label="docker-image://docker.io/docker/dockerfile-copy:v0.1.9" shape="ellipse"];
  "sha256:73f494e3e5baaa46b1a4cc3a4d1c59c049e82cd5cd70b3c166fe09cbe4cb143e" [label="local://context" shape="ellipse"];
  "sha256:8f440bbee7e64fd9a1846d02a7e195458de7f91994244c95683866804fed65d6" [label="docker-image://docker.io/library/golang:1.12" shape="ellipse"];
  "sha256:1d9bc5098154416cf2d5ba0b0aaba6ab88348e4cc0a728bf0f37a7db32c36426" [label="copy --unpack /src-0 go" shape="box"];
  "sha256:0f1183cf8ee0399b25b89d75d371ff7dd2cf4b9f10c73a0761ecfd29ef4a9120" [label="/bin/sh -c go build -o stage1_bin" shape="box"];
  "sha256:348c2dded9e336b03cea79e2cfebbe0f4a7189cf306fa283380ed4b7e5e51d32" [label="/bin/sh -c go build -o stage0_bin" shape="box"];
  "sha256:1b2db4270fb36ed9837304839bbaa1d532bd3893fbd2bfd57863d6dbe85e0e7c" [label="copy /src-0/stage0_bin ./" shape="box"];
  "sha256:2026e0c8c202590a664fc4e03de5d236045bea576d68fd7078d76e86cdecf5ba" [label="copy /src-0/stage1_bin ./" shape="box"];
  "sha256:4b0a5cf7b4ce98d7862bf8cbb594ce3527717bedb8675dc37dc0adcd4512d1f9" [label="sha256:4b0a5cf7b4ce98d7862bf8cbb594ce3527717bedb8675dc37dc0adcd4512d1f9" shape="plaintext"];
  "sha256:1c5320070ff30eecf3265f227b2646e54428945092b0866a55da4bb20415f066" -> "sha256:1d9bc5098154416cf2d5ba0b0aaba6ab88348e4cc0a728bf0f37a7db32c36426" [label=""];
  "sha256:8f440bbee7e64fd9a1846d02a7e195458de7f91994244c95683866804fed65d6" -> "sha256:1d9bc5098154416cf2d5ba0b0aaba6ab88348e4cc0a728bf0f37a7db32c36426" [label="/dest"];
  "sha256:73f494e3e5baaa46b1a4cc3a4d1c59c049e82cd5cd70b3c166fe09cbe4cb143e" -> "sha256:1d9bc5098154416cf2d5ba0b0aaba6ab88348e4cc0a728bf0f37a7db32c36426" [label="/src-0"];
  "sha256:1d9bc5098154416cf2d5ba0b0aaba6ab88348e4cc0a728bf0f37a7db32c36426" -> "sha256:0f1183cf8ee0399b25b89d75d371ff7dd2cf4b9f10c73a0761ecfd29ef4a9120" [label=""];
  "sha256:1d9bc5098154416cf2d5ba0b0aaba6ab88348e4cc0a728bf0f37a7db32c36426" -> "sha256:348c2dded9e336b03cea79e2cfebbe0f4a7189cf306fa283380ed4b7e5e51d32" [label=""];
  "sha256:1c5320070ff30eecf3265f227b2646e54428945092b0866a55da4bb20415f066" -> "sha256:1b2db4270fb36ed9837304839bbaa1d532bd3893fbd2bfd57863d6dbe85e0e7c" [label=""];
  "sha256:8f440bbee7e64fd9a1846d02a7e195458de7f91994244c95683866804fed65d6" -> "sha256:1b2db4270fb36ed9837304839bbaa1d532bd3893fbd2bfd57863d6dbe85e0e7c" [label="/dest"];
  "sha256:348c2dded9e336b03cea79e2cfebbe0f4a7189cf306fa283380ed4b7e5e51d32" -> "sha256:1b2db4270fb36ed9837304839bbaa1d532bd3893fbd2bfd57863d6dbe85e0e7c" [label="/src-0/stage0_bin"];
  "sha256:1c5320070ff30eecf3265f227b2646e54428945092b0866a55da4bb20415f066" -> "sha256:2026e0c8c202590a664fc4e03de5d236045bea576d68fd7078d76e86cdecf5ba" [label=""];
  "sha256:1b2db4270fb36ed9837304839bbaa1d532bd3893fbd2bfd57863d6dbe85e0e7c" -> "sha256:2026e0c8c202590a664fc4e03de5d236045bea576d68fd7078d76e86cdecf5ba" [label="/dest"];
  "sha256:0f1183cf8ee0399b25b89d75d371ff7dd2cf4b9f10c73a0761ecfd29ef4a9120" -> "sha256:2026e0c8c202590a664fc4e03de5d236045bea576d68fd7078d76e86cdecf5ba" [label="/src-0/stage1_bin"];
  "sha256:2026e0c8c202590a664fc4e03de5d236045bea576d68fd7078d76e86cdecf5ba" -> "sha256:4b0a5cf7b4ce98d7862bf8cbb594ce3527717bedb8675dc37dc0adcd4512d1f9" [label=""];
}`,
			),
		},
	}

	for _, tt := range tests {
		got, err := docker2dot.Docker2Dot(tt.input)
		if err != nil {
			t.Errorf("got unexpected error: %+v", err)
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got unexpected result:\ngot: %+v\nwant: %+v\n", string(got), string(tt.want))
		}
	}
}
