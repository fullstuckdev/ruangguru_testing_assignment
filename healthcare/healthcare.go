package main

import (
    "fmt"
    "time"
)

type Reservation struct {
    Name     string
    DateTime time.Time
}

type Doctor struct {
    ID        string
    Name      string
    Specialty string
}

var doctors = make(map[string]Doctor)

func MakeReservation(name string, dateTime time.Time, reservations map[string]Reservation) {
    reservations[name] = Reservation{name, dateTime}
    fmt.Println("Reservasi berhasil untuk", name, "pada tanggal", dateTime.Format("02/01/2006 15:04"))
}

func ShowReservations(reservations map[string]Reservation) {
    fmt.Println("Daftar Reservasi:")
    for _, reservation := range reservations {
        fmt.Println("Nama:", reservation.Name, "| Tanggal:", reservation.DateTime.Format("02/01/2006 15:04"))
    }
}

func CreateDoctor(id, name, specialty string) {
	if doctors == nil {
		doctors = make(map[string]Doctor)
	}
	doctors[id] = Doctor{ID: id, Name: name, Specialty: specialty}
}

func UpdateDoctor(id string, name string, specialty string) {
	if _, ok := doctors[id]; ok {
		doctors[id] = Doctor{ID: id, Name: name, Specialty: specialty}
		fmt.Println("Data dokter dengan ID", id, "berhasil diperbarui.")
	} else {
		fmt.Println("Dokter dengan ID", id, "tidak ditemukan.")
	}
}

func ReadDoctor(id string) {
    if doctor, ok := doctors[id]; ok {
        fmt.Println("ID:", doctor.ID, "| Nama:", doctor.Name, "| Spesialis:", doctor.Specialty)
    } else {
        fmt.Println("Dokter dengan ID", id, "tidak ditemukan.")
    }
}


func DeleteDoctor(id string) {
    if _, ok := doctors[id]; ok {
        delete(doctors, id)
        fmt.Println("Dokter dengan ID", id, "berhasil dihapus.")
    } else {
        fmt.Println("Dokter dengan ID", id, "tidak ditemukan.")
    }
}


func main() {
    reservations := make(map[string]Reservation)

    fmt.Println("Selamat datang di Sistem Reservasi Rumah Sakit")
    fmt.Println("1. Buat Reservasi")
    fmt.Println("2. Tampilkan Semua Reservasi")
    fmt.Println("3. Tambah Dokter")
    fmt.Println("4. Lihat Dokter")
    fmt.Println("5. Perbarui Dokter")
    fmt.Println("6. Hapus Dokter")
    fmt.Println("7. Keluar")

    for {
        var choice int
        fmt.Print("Pilih menu: ")
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            var name string
            var dateStr string
            var timeStr string

            fmt.Print("Masukkan nama Anda: ")
            fmt.Scanln(&name)

            fmt.Print("Masukkan tanggal reservasi (format: dd/mm/yyyy): ")
            fmt.Scanln(&dateStr)

            fmt.Print("Masukkan jam reservasi (format: hh:mm): ")
            fmt.Scanln(&timeStr)

            dateTimeStr := fmt.Sprintf("%s %s", dateStr, timeStr)
            dateTime, err := time.Parse("02/01/2006 15:04", dateTimeStr)
            if err != nil {
                fmt.Println("Format tanggal dan waktu salah. Coba lagi.")
                continue
            }

            MakeReservation(name, dateTime, reservations)

        case 2:
            ShowReservations(reservations)

        case 3:
            var id string
            var name string
            var specialty string

            fmt.Print("Masukkan ID dokter: ")
            fmt.Scanln(&id)

            fmt.Print("Masukkan nama dokter: ")
            fmt.Scanln(&name)

            fmt.Print("Masukkan spesialis dokter: ")
            fmt.Scanln(&specialty)

            CreateDoctor(id, name, specialty)

        case 4:
            var id string

            fmt.Print("Masukkan ID dokter: ")
            fmt.Scanln(&id)

            ReadDoctor(id)

        case 5:
            var id string
            var name string
            var specialty string

            fmt.Print("Masukkan ID dokter: ")
            fmt.Scanln(&id)

            fmt.Print("Masukkan nama baru dokter: ")
            fmt.Scanln(&name)

            fmt.Print("Masukkan spesialis baru dokter: ")
            fmt.Scanln(&specialty)

            UpdateDoctor(id, name, specialty)

        case 6:
            var id string

            fmt.Print("Masukkan ID dokter: ")
            fmt.Scanln(&id)

            DeleteDoctor(id)
        
        case 7:
            fmt.Println("Terima kasih telah menggunakan layanan kami.")
            return

        default:
            fmt.Println("Pilihan tidak valid. Silakan pilih 1, 2, 3, 4, 5, 6, dan 7.")
        }
    }
}
