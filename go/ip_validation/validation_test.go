package main

import "testing"

func TestValidation(t *testing.T) {
	if ip_validation("invalid_ip") != false {
		t.Errorf("Validation failed on invalid string")
	}

	if ip_validation("12.255.56.1") != true {
		t.Errorf("Validation failed on valid IPV4 address")
	}

	if ip_validation("FE80:CD00:0000:0CDE:1257:0000:211E:729C") != false {
		t.Errorf("Validation failed on IPV6 test")
	}

	if ip_validation("123.045.067.089") != false {
		t.Errorf("Validation failed on lead 0 IPV4")
	}
}
