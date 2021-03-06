// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"io/ioutil"
	"testing"

	"github.com/GoogleCloudPlatform/golang-samples/internal/testutil"

	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
	speech "google.golang.org/genproto/googleapis/cloud/speech/v1beta1"
)

func TestRecognize(t *testing.T) {
	testutil.SystemTest(t)

	ctx := context.Background()
	conn, err := transport.DialGRPC(ctx,
		option.WithEndpoint("speech.googleapis.com:443"),
		option.WithScopes("https://www.googleapis.com/auth/cloud-platform"),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	c := speech.NewSpeechClient(conn)

	data, err := ioutil.ReadFile("./quit.raw")
	if err != nil {
		t.Fatal(err)
	}

	rresp, err := recognize(ctx, c, &data)
	if err != nil {
		t.Fatal(err)
	}
	if len(rresp.Results) < 1 {
		t.Fatal("want recognize results; got none")
	}

	result := rresp.Results[0]
	if len(result.Alternatives) < 1 {
		t.Fatal("got no alternatives; want at least one")
	}
	if got, want := result.Alternatives[0].Transcript, "quit"; got != want {
		t.Errorf("got transcript: %q; want %q", got, want)
	}
}
