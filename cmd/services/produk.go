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
	
	var produk []*produkPb.Product

	rows, err := p.DB.Table("produk AS p").
		Joins("LEFT JOIN ulasan AS u ON u.id_ulasan = p.ulasan_id").
		Select("p.id","p.namaproduk","p.harga","p.kategori","p.stok","u.id_ulasan").
		Rows()

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var produk *produkPb.Product
		var ulasan *produkPb.Ulasan

		if rows.Scan(&produk.Id, &produk.NamaProduk, &produk.Harga, &produk.Kategori, &produk.Stok, &ulasan.IdUlasan); err != nil {
			log.Fatalf("Gagal mengambil row data %v",err.Error())
		}
		produk.Ulasan = ulasan
		produk = append(produk, &produk)
	}

	response := &produkPb.Product{
		Ulasan: []*produkPb.Ulasan{},
	}
	return response, nil
}