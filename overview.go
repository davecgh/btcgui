/*
 * Copyright (c) 2013 Conformal Systems LLC <info@conformal.com>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package main

import (
	"fmt"
	"github.com/conformal/gotk3/gtk"
	"log"
	"math"
	"strconv"
	"time"
)

type txDirection int

const (
	SEND txDirection = iota
	RECV
)

var (
	Overview = struct {
		Balance       *gtk.Label
		Unconfirmed   *gtk.Label
		NTransactions *gtk.Label
	}{}
)

func createWalletInfo() *gtk.Widget {
	grid, err := gtk.GridNew()
	if err != nil {
		log.Fatal(err)
	}

	header, err := gtk.LabelNew("")
	if err != nil {
		log.Fatal(err)
	}
	header.SetMarkup("<b>Wallet</b>")
	header.OverrideFont("sans-serif 16")
	header.SetHAlign(gtk.ALIGN_START)
	grid.Attach(header, 0, 0, 1, 1)

	balance, err := gtk.LabelNew("Balance:")
	if err != nil {
		log.Fatal(err)
	}
	balance.SetHAlign(gtk.ALIGN_START)
	grid.Attach(balance, 0, 1, 1, 1)

	unconfirmed, err := gtk.LabelNew("Unconfirmed:")
	if err != nil {
		log.Fatal(err)
	}
	unconfirmed.SetHAlign(gtk.ALIGN_START)
	grid.Attach(unconfirmed, 0, 2, 1, 1)

	transactions, err := gtk.LabelNew("Number of transactions:")
	if err != nil {
		log.Fatal(err)
	}
	transactions.SetHAlign(gtk.ALIGN_START)
	grid.Attach(transactions, 0, 3, 1, 1)

	balance, err = gtk.LabelNew("")
	if err != nil {
		log.Fatal(err)
	}
	balance.SetHAlign(gtk.ALIGN_START)
	grid.Attach(balance, 1, 1, 1, 1)
	Overview.Balance = balance

	unconfirmed, err = gtk.LabelNew("")
	if err != nil {
		log.Fatal(err)
	}
	grid.Attach(unconfirmed, 1, 2, 1, 1)
	Overview.Unconfirmed = unconfirmed

	transactions, err = gtk.LabelNew(strconv.Itoa(2))
	if err != nil {
		log.Fatal(err)
	}
	transactions.SetHAlign(gtk.ALIGN_START)
	grid.Attach(transactions, 1, 3, 1, 1)
	Overview.NTransactions = transactions

	return &grid.Container.Widget
}

func createTxInfo() *gtk.Widget {
	grid, err := gtk.GridNew()
	if err != nil {
		log.Fatal(err)
	}
	grid.SetOrientation(gtk.ORIENTATION_VERTICAL)

	l, err := gtk.LabelNew("")
	if err != nil {
		log.Fatal(err)
	}
	l.SetMarkup("<b>Recent Transactions</b>")
	l.OverrideFont("sans-serif 10")
	l.SetHAlign(gtk.ALIGN_START)
	grid.Add(l)

	// TODO(jrick): connect this
	grid.Add(createTxLabel(SEND, 1.0, "1234567890", time.Now()))
	grid.Add(createTxLabel(RECV, 1.0, "0987654321", time.Now()))

	return &grid.Container.Widget
}

func createTxLabel(dir txDirection, amt float64, addr string, t time.Time) *gtk.Widget {
	grid, err := gtk.GridNew()
	if err != nil {
		log.Fatal(err)
	}
	grid.SetHExpand(true)

	var amtLabel *gtk.Label
	var description *gtk.Label
	var icon *gtk.Image
	switch dir {
	case SEND:
		s := "-" +
			strconv.FormatFloat(math.Abs(0), 'f', 8, 64) +
			" BTC"

		amtLabel, err = gtk.LabelNew(s)
		if err != nil {
			log.Fatal(err)
		}

		description, err = gtk.LabelNew(fmt.Sprintf("Purchase (%s)", addr))
		if err != nil {
			log.Fatal(err)
		}

		icon, err = gtk.ImageNewFromStock(gtk.STOCK_GO_FORWARD,
			gtk.ICON_SIZE_SMALL_TOOLBAR)
		if err != nil {
			log.Fatal(err)
		}
	case RECV:
		s := strconv.FormatFloat(math.Abs(0), 'f', 8, 64) +
			" BTC"

		amtLabel, err = gtk.LabelNew(s)
		if err != nil {
			log.Fatal(err)
		}

		description, err = gtk.LabelNew(fmt.Sprintf("Payment (%s)", addr))
		if err != nil {
			log.Fatal(err)
		}

		icon, err = gtk.ImageNewFromStock(gtk.STOCK_GO_BACK,
			gtk.ICON_SIZE_SMALL_TOOLBAR)
		if err != nil {
			log.Fatal(err)
		}
	}
	grid.Attach(icon, 0, 0, 2, 2)
	grid.Attach(description, 2, 1, 2, 1)
	description.SetHAlign(gtk.ALIGN_START)
	grid.Attach(amtLabel, 3, 0, 1, 1)
	amtLabel.SetHAlign(gtk.ALIGN_END)
	amtLabel.SetHExpand(true)

	date, err := gtk.LabelNew(t.Format("Jan 2, 2006 at 3:04 PM"))
	if err != nil {
		log.Fatal(err)
	}
	grid.Attach(date, 2, 0, 1, 1)
	date.SetHAlign(gtk.ALIGN_START)

	grid.SetHAlign(gtk.ALIGN_FILL)

	return &grid.Container.Widget
}

func createOverview() *gtk.Widget {
	grid, err := gtk.GridNew()
	if err != nil {
		log.Fatal(err)
	}

	grid.SetColumnHomogeneous(true)
	grid.Add(createWalletInfo())
	grid.Add(createTxInfo())

	return &grid.Container.Widget
}