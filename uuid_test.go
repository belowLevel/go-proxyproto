package proxyproto

import "testing"

func BenchmarkUUID(b *testing.B) {
	for b.Loop() {
		NewUUIDV4()
		//NewUUIDV1()
		//NewUUIDV6()
		//NewUUIDV7()
	}

}
