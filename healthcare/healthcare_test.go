package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"Testing/healthcare"
	"time"
	"os"
	"strings"
	"io/ioutil"
)

var _ = Describe("Healthcare System", func() {
	var (
		reservations map[string]main.Reservation
		doctors      map[string]main.Doctor
	)

	BeforeEach(func() {
		reservations = make(map[string]main.Reservation)
		doctors = make(map[string]main.Doctor)
	})

	Describe("Reservation Management", func() {
		Context("MakeReservation", func() {
			It("should add a reservation to the map", func() {
				name := "Taufik Mulyawan"
				dateTime := time.Now()
				main.MakeReservation(name, dateTime, reservations)
				Expect(reservations).To(HaveLen(1))
			})
		})

		Context("ShowReservations", func() {
			It("should display all reservations", func() {
				name := "Shila Afifah"
				dateTime := time.Now().Add(24 * time.Hour)
				main.MakeReservation(name, dateTime, reservations)

				oldStdout := os.Stdout
				r, w, _ := os.Pipe()
				os.Stdout = w

				main.ShowReservations(reservations)

				w.Close()
				out, _ := ioutil.ReadAll(r)
				os.Stdout = oldStdout

				Expect(strings.Contains(string(out), name)).To(BeTrue())
			})
		})
	})

	Describe("Doctor Management", func() {

		Context("ReadDoctor", func() {
			It("should display doctor information", func() {
				id := "1"
				name := "Dr. Taufik"
				specialty := "Cardiology"
				main.CreateDoctor(id, name, specialty)

				oldStdout := os.Stdout
				r, w, _ := os.Pipe()
				os.Stdout = w

				main.ReadDoctor(id)

				w.Close()
				out, _ := ioutil.ReadAll(r)
				os.Stdout = oldStdout

				Expect(strings.Contains(string(out), id)).To(BeTrue())
				Expect(strings.Contains(string(out), name)).To(BeTrue())
				Expect(strings.Contains(string(out), specialty)).To(BeTrue())
			})
		})


		Context("DeleteDoctor", func() {
			It("should delete a doctor from the map", func() {
				id := "1"
				name := "Dr. Taufik"
				specialty := "Cardiology"
				main.CreateDoctor(id, name, specialty)
				main.DeleteDoctor(id)
				Expect(doctors).ToNot(HaveKey(id))
			})

			It("should not delete a doctor if it does not exist", func() {
				id := "100"
				main.DeleteDoctor(id)
				Expect(doctors).ToNot(HaveKey(id))
			})
		})
	})
})