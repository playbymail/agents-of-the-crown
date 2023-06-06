// Copyright (c) 2023 Michael D Henderson. All rights reserved.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"io"
	"math/rand"
)

type Deck []*Card

func (d Deck) Dump(w io.Writer, cardsPerRow int) {
	cardsPrinted := 0
	for _, card := range d {
		_, _ = fmt.Fprintf(w, "  %3s", card.String())
		if cardsPrinted = cardsPrinted + 1; cardsPrinted == cardsPerRow {
			_, _ = fmt.Fprintln(w, "")
			cardsPrinted = 0
		}
	}
	if cardsPrinted != 0 {
		_, _ = fmt.Fprintln(w, "")
	}
}

func (d Deck) Shuffle() {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

func fullDeck() Deck {
	return []*Card{
		{Suit: SPADES, Rank: ACE, Face: "🂡"}, {Suit: HEARTS, Rank: ACE, Face: "🂱"}, {Suit: DIAMONDS, Rank: ACE, Face: "🃁"}, {Suit: CLUBS, Rank: ACE, Face: "🃑"},
		{Suit: SPADES, Rank: TWO, Face: "🂢"}, {Suit: HEARTS, Rank: TWO, Face: "🂲"}, {Suit: DIAMONDS, Rank: TWO, Face: "🃂"}, {Suit: CLUBS, Rank: TWO, Face: "🃒"},
		{Suit: SPADES, Rank: THREE, Face: "🂣"}, {Suit: HEARTS, Rank: THREE, Face: "🂳"}, {Suit: DIAMONDS, Rank: THREE, Face: "🃃"}, {Suit: CLUBS, Rank: THREE, Face: "🃓"},
		{Suit: SPADES, Rank: FOUR, Face: "🂤"}, {Suit: HEARTS, Rank: FOUR, Face: "🂴"}, {Suit: DIAMONDS, Rank: FOUR, Face: "🃄"}, {Suit: CLUBS, Rank: FOUR, Face: "🃔"},
		{Suit: SPADES, Rank: FIVE, Face: "🂥"}, {Suit: HEARTS, Rank: FIVE, Face: "🂵"}, {Suit: DIAMONDS, Rank: FIVE, Face: "🃅"}, {Suit: CLUBS, Rank: FIVE, Face: "🃕"},
		{Suit: SPADES, Rank: SIX, Face: "🂦"}, {Suit: HEARTS, Rank: SIX, Face: "🂶"}, {Suit: DIAMONDS, Rank: SIX, Face: "🃆"}, {Suit: CLUBS, Rank: SIX, Face: "🃖"},
		{Suit: SPADES, Rank: SEVEN, Face: "🂧"}, {Suit: HEARTS, Rank: SEVEN, Face: "🂷"}, {Suit: DIAMONDS, Rank: SEVEN, Face: "🃇"}, {Suit: CLUBS, Rank: SEVEN, Face: "🃗"},
		{Suit: SPADES, Rank: EIGHT, Face: "🂨"}, {Suit: HEARTS, Rank: EIGHT, Face: "🂸"}, {Suit: DIAMONDS, Rank: EIGHT, Face: "🃈"}, {Suit: CLUBS, Rank: EIGHT, Face: "🃘"},
		{Suit: SPADES, Rank: NINE, Face: "🂩"}, {Suit: HEARTS, Rank: NINE, Face: "🂹"}, {Suit: DIAMONDS, Rank: NINE, Face: "🃉"}, {Suit: CLUBS, Rank: NINE, Face: "🃙"},
		{Suit: SPADES, Rank: TEN, Face: "🂪"}, {Suit: HEARTS, Rank: TEN, Face: "🂺"}, {Suit: DIAMONDS, Rank: TEN, Face: "🃊"}, {Suit: CLUBS, Rank: TEN, Face: "🃚"},
		{Suit: SPADES, Rank: JACK, Face: "🂫"}, {Suit: HEARTS, Rank: JACK, Face: "🂻"}, {Suit: DIAMONDS, Rank: JACK, Face: "🃋"}, {Suit: CLUBS, Rank: JACK, Face: "🃛"},
		{Suit: SPADES, Rank: KNIGHT, Face: "🂬"}, {Suit: HEARTS, Rank: KNIGHT, Face: "🂼"}, {Suit: DIAMONDS, Rank: KNIGHT, Face: "🃌"}, {Suit: CLUBS, Rank: KNIGHT, Face: "🃜"},
		{Suit: SPADES, Rank: QUEEN, Face: "🂭"}, {Suit: HEARTS, Rank: QUEEN, Face: "🂽"}, {Suit: DIAMONDS, Rank: QUEEN, Face: "🃍"}, {Suit: CLUBS, Rank: QUEEN, Face: "🃝"},
		{Suit: SPADES, Rank: KING, Face: "🂮"}, {Suit: HEARTS, Rank: KING, Face: "🂾"}, {Suit: DIAMONDS, Rank: KING, Face: "🃎"}, {Suit: CLUBS, Rank: KING, Face: "🃞"},
		{Rank: JOKER, Face: "🂿"}, // red joker
		{Rank: JOKER, Face: "🃟"}, // white joker
	}
}

func newDeck(accept func(Card) bool) Deck {
	var deck []*Card
	for _, card := range fullDeck() {
		if accept(*card) {
			deck = append(deck, card)
		}
	}

	for _, card := range deck {
		card.Back = "🂠"
	}

	return deck
}

type Card struct {
	Suit Suit
	Rank Rank
	Face string
	Back string
}

type Rank int

const (
	JOKER Rank = iota
	ACE
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	KNIGHT
	QUEEN
	KING
)

func (r Rank) String() string {
	switch r {
	case JOKER:
		return "j"
	case ACE:
		return "A"
	case TWO:
		return "2"
	case THREE:
		return "3"
	case FOUR:
		return "4"
	case FIVE:
		return "5"
	case SIX:
		return "6"
	case SEVEN:
		return "7"
	case EIGHT:
		return "9"
	case NINE:
		return "9"
	case TEN:
		return "10"
	case JACK:
		return "J"
	case KNIGHT:
		return "k"
	case QUEEN:
		return "Q"
	case KING:
		return "K"
	}
	panic(fmt.Sprintf("assert(rank != %d)", int(r)))
}

type Suit int

const (
	CLUBS Suit = iota
	DIAMONDS
	HEARTS
	SPADES
)

func (s Suit) String() string {
	switch s {
	case CLUBS:
		return "♧"
	case DIAMONDS:
		return "♢"
	case HEARTS:
		return "♡"
	case SPADES:
		return "♤"
	}
	panic(fmt.Sprintf("assert(suit != %d)", int(s)))
}

func (c Card) IsBanner() bool {
	return c.Rank == ACE
}

func (c Card) IsCommon() bool {
	return TWO <= c.Rank && c.Rank <= NINE
}

func (c Card) IsNoble() bool {
	return c.Rank == TEN || c.Rank == JACK || c.Rank == QUEEN || c.Rank == KING
}

func (c Card) String() string {
	return fmt.Sprintf("%s%s", c.Rank, c.Suit)
}
