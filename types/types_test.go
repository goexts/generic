package types_test

import (
	"fmt"
	"testing"

	"github.com/goexts/generic/object"
	"github.com/goexts/generic/types"
)

func TestBasicTypes(t *testing.T) {
	// Test String
	s1 := types.NewString("hello")
	s2 := types.NewString("hello")
	s3 := types.NewString("world")

	if s1.String() != "hello" {
		t.Errorf("String.String() failed, got: %s, want: %s", s1.String(), "hello")
	}
	if !s1.Equals(s2) {
		t.Errorf("String.Equals() failed, s1 should be equal to s2")
	}
	if s1.Equals(s3) {
		t.Errorf("String.Equals() failed, s1 should not be equal to s3")
	}
	if s1.HashCode() != s2.HashCode() {
		t.Errorf("String.HashCode() failed, s1 and s2 should have the same hash code")
	}
	if s1.HashCode() == s3.HashCode() {
		t.Errorf("String.HashCode() failed, s1 and s3 should have different hash codes")
	}

	// Test Int
	i1 := types.NewInt(100)
	i2 := types.NewInt(100)
	i3 := types.NewInt(200)

	if i1.String() != "100" {
		t.Errorf("Int.String() failed, got: %s, want: %s", i1.String(), "100")
	}
	if !i1.Equals(i2) {
		t.Errorf("Int.Equals() failed, i1 should be equal to i2")
	}
	if i1.Equals(i3) {
		t.Errorf("Int.Equals() failed, i1 should not be equal to i3")
	}
	if i1.HashCode() != i2.HashCode() {
		t.Errorf("Int.HashCode() failed, i1 and i2 should have the same hash code")
	}

	// Test Boolean
	b1 := types.NewBoolean(true)
	b2 := types.NewBoolean(true)
	b3 := types.NewBoolean(false)

	if b1.String() != "true" {
		t.Errorf("Boolean.String() failed, got: %s, want: %s", b1.String(), "true")
	}
	if !b1.Equals(b2) {
		t.Errorf("Boolean.Equals() failed, b1 should be equal to b2")
	}
	if b1.Equals(b3) {
		t.Errorf("Boolean.Equals() failed, b1 should not be equal to b3")
	}
	if b1.HashCode() != b2.HashCode() {
		t.Errorf("Boolean.HashCode() failed, b1 and b2 should have the same hash code")
	}

	// Test Bytes
	bytes1 := types.NewBytes([]byte("hello"))
	bytes2 := types.NewBytes([]byte("hello"))
	bytes3 := types.NewBytes([]byte("world"))

	if bytes1.String() != "hello" {
		t.Errorf("Bytes.String() failed, got: %s, want: %s", bytes1.String(), "hello")
	}
	if !bytes1.Equals(bytes2) {
		t.Errorf("Bytes.Equals() failed, bytes1 should be equal to bytes2")
	}
	if bytes1.Equals(bytes3) {
		t.Errorf("Bytes.Equals() failed, bytes1 should not be equal to bytes3")
	}
	if bytes1.HashCode() != bytes2.HashCode() {
		t.Errorf("Bytes.HashCode() failed, bytes1 and bytes2 should have the same hash code")
	}
	if bytes1.HashCode() == bytes3.HashCode() {
		t.Errorf("Bytes.HashCode() failed, bytes1 and bytes3 should have different hash codes")
	}

	// Test Runes
	runes1 := types.NewRunes([]rune("hello"))
	runes2 := types.NewRunes([]rune("hello"))
	runes3 := types.NewRunes([]rune("world"))

	if runes1.String() != "hello" {
		t.Errorf("Runes.String() failed, got: %s, want: %s", runes1.String(), "hello")
	}
	if !runes1.Equals(runes2) {
		t.Errorf("Runes.Equals() failed, runes1 should be equal to runes2")
	}
	if runes1.Equals(runes3) {
		t.Errorf("Runes.Equals() failed, runes1 should not be equal to runes3")
	}
	if runes1.HashCode() != runes2.HashCode() {
		t.Errorf("Runes.HashCode() failed, runes1 and runes2 should have the same hash code")
	}
	if runes1.HashCode() == runes3.HashCode() {
		t.Errorf("Runes.HashCode() failed, runes1 and runes3 should have different hash codes")
	}

	// Test Polymorphism
	var obj object.Object = types.NewInt(42)
	fmt.Printf("Polymorphic object: %s, HashCode: %d\n", obj.String(), obj.HashCode())

	obj = types.NewString("polymorphism")
	fmt.Printf("Polymorphic object: %s, HashCode: %d\n", obj.String(), obj.HashCode())
}
