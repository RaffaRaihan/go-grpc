package services

import (
	"context"
	produkPb "go-grpc/pb/produk"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type ProdukService struct{
	produkPb.UnimplementedProductClientServer
	DB *gorm.DB
}

func (p *ProdukService) GetProduks(context.Context, *produkPb.Empty) (*produkPb.Product, error) {
	
	var produk produkPb.Product

	rows, err := p.DB.Table("spesifikasi AS s","ulasan AS u").
		Joins("LEFT JOIN produk AS p ON p.id = s.produk_id").
		Joins("LEFT JOIN produk AS p ON p.id = u.produk_id").
		Select("p.id", "p.nama_produk", "p.kategori", "p.harga", "p.stok", "p.deskripsi", "s.id", "p.merek", "p.tanggal_rilis", "p.rating", "u.id_ulasan").
		Rows()

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var product produkPb.Product
		var spesifikasi produkPb.Spesifikasi
		var ulasan produkPb.Ulasan

		if err := rows.Scan(product.Id, product.NamaProduk, product.Kategori, product.Harga, product.Stok, product.Deskripsi, spesifikasi.IdSpesifikasi, product.Merek, product.TanggalRilis, product.Rating, ulasan.IdUlasan); err != nil{
			log.Fatalf("Gagal mengambil data")
		}
		product.Spesifikasi = &spesifikasi
		product.Ulasan = &ulasan
		product = append(product, &produk)
	}

	response := &produkPb.Product{
		Id: produk.Id,
		NamaProduk: produk.NamaProduk,
		Kategori: produk.Kategori,
		Harga: produk.Harga,
		Stok: produk.Stok,
		Deskripsi: produk.Deskripsi,
		Spesifikasi: produk.Spesifikasi,
		Merek: produk.Merek,
		TanggalRilis: produk.TanggalRilis,
		Rating: produk.Rating,
		Ulasan: produk.Ulasan,
	}

	return response, nil
}