package main

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	command := "insert into karyawan(nama, role, alamat) values('Alam', 'SWE', 'Jalan Kembangan Utara kontrakan Babeh Roy')"

	_, err := db.ExecContext(ctx, command)
	if err != nil {
		panic(err)
	}

	fmt.Println("Berhasil menambahkan data")
}

func TestQuery(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	command := "select nama, role, alamat from karyawan"

	rows, err := db.QueryContext(ctx, command)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var nama, role, alamat string
		err := rows.Scan(&nama, &role, &alamat)
		if err != nil {
			panic(err)
		}
		fmt.Println("nama: ", nama)
		fmt.Println("role: ", role)
		fmt.Println("alamat: ", alamat)
	}
}
