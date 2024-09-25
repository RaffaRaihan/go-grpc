package services

import (
	"context"
	produkPb "go-grpc/pb/produk"
)

type ProdukService struct{
	produkPb.UnimplementedProdukServiceServer
}

func (p *ProdukService) GetProduk(context.Context, *produkPb.Empty) (*produkPb.Product, error) {
	produk := *&produkPb.Product{
		Id: 101,
		NamaProduk: "Laptop Gaming",
		Kategori: "Elektronik",
		Harga: 1500000,
		Stok: 25,
		Deskripsi: "Laptop gaming dengan spesifikasi tinggi untuk pengalaman bermain game yang luar biasa.",
		Spesifikasi: &produkPb.Spesifikasi{
			Procesor: "Intel Core i7",
			Ram: "16GB",
			Penyimpanan: "1TB SSD",
			Layar: "15.6 inci",
			Grafis: "NDIVIA RTX 3060",
		},
		Merek: "BrandX",
		TanggalRilis: "2023-07-15",
		Rating: 4.5,
		Ulasan: []*produkPb.Ulasan{
			{
				IdUlasan: 1,
				NamaPelanggan: "Budi",
				TanggalUlasan: "2023-07-20",
				Komentar: "Laptopnya sangat cepat dan lancar untuk bermain game.",
				Rating: 5,
			},
			{
				IdUlasan: 2,
				NamaPelanggan: "Siti",
				TanggalUlasan: "2023-07-22",
				Komentar: "Desainnya keren tapi agak panas kalau dipakai lama.",
				Rating: 4,
			},
		},


	}
	return &produk, nil
}