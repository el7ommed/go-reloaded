package main

import (
	"testing"

	"learn.reboot01.com/git/moadwan/go-reloaded/autocorrect"
)

func TestModifyLine001(t *testing.T) {
	got := autocorrect.ModifyLine("If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?")
	want := "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestModifyLine002(t *testing.T) {
	got := autocorrect.ModifyLine("I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure")
	want := "I have to pack 5 outfits. Packed 26 just to be sure"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestModifyLine003(t *testing.T) {
	got := autocorrect.ModifyLine("Don not be sad ,because sad backwards is das . And das not good")
	want := "Don not be sad, because sad backwards is das. And das not good"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestModifyLine004(t *testing.T) {
	got := autocorrect.ModifyLine("harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '")
	want := "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestModifyLine1(t *testing.T) {
	got := autocorrect.ModifyLine("it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.")
	want := "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair."

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestModifyLine2(t *testing.T) {
	got := autocorrect.ModifyLine("Simply add 42 (hex) and 10 (bin) and you will see the result is 68.")
	want := "Simply add 66 and 2 and you will see the result is 68."

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestModifyLine3(t *testing.T) {
	got := autocorrect.ModifyLine("There is no greater agony than bearing a untold story inside you.")
	want := "There is no greater agony than bearing an untold story inside you."

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestModifyLine4(t *testing.T) {
	got := autocorrect.ModifyLine("Punctuation tests are ... kinda boring ,don't you think !?")
	want := "Punctuation tests are... kinda boring, don't you think!?"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestModifyLine5(t *testing.T) {
	got := autocorrect.ModifyLine("1E (hex) files were added")
	want := "30 files were added"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestModifyLine6(t *testing.T) {
	got := autocorrect.ModifyLine("It has been 10 (bin) years")
	want := "It has been 2 years"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestModifyLine7(t *testing.T) {
	got := autocorrect.ModifyLine("Ready, set, go (up) !")
	want := "Ready, set, GO!"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestModifyLine8(t *testing.T) {
	got := autocorrect.ModifyLine("I should stop SHOUTING (low)")
	want := "I should stop shouting"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestModifyLine9(t *testing.T) {
	got := autocorrect.ModifyLine("Welcome to the Brooklyn bridge (cap)")
	want := "Welcome to the Brooklyn Bridge"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestModifyLine10(t *testing.T) {
	got := autocorrect.ModifyLine("This is so exciting (up, 2)")
	want := "This is SO EXCITING"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
