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
			return nil, status.Error(codes.NotFound, err.Error())
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

func (p *ProdukService) GetProduk(ctx context.Context, id *produkPb.Id) (*produkPb.Product, error) {
	row := p.DB.Table("spesifikasi AS s","ulasan AS u").
		Joins("LEFT JOIN produk AS p ON p.id = s.produk_id").
		Joins("LEFT JOIN produk AS p ON p.id = u.produk_id").
		Select("p.id", "p.nama_produk", "p.kategori", "p.harga", "p.stok", "p.deskripsi", "s.id", "p.merek", "p.tanggal_rilis", "p.rating", "u.id_ulasan").
		Where("p.id = ?", id.GetId()).
		Row()

		var product produkPb.Product

		if err := row.Scan(product.Id, product.NamaProduk, product.Kategori, product.Harga, product.Stok, product.Deskripsi, spesifikasi.IdSpesifikasi, product.Merek, product.TanggalRilis, product.Rating, ulasan.IdUlasan); err != nil{
			log.Fatalf("Gagal mengambil data")
		}
		product.Id = &product
		
		return &product, nil 
}

func (p *ProdukService) CreateProduk(ctx context.Context, produkData *produkPb.Product) (*produkPb.Id, error) {

	var Response produkPb.Id

	err := p.DB.Transaction(func(tx *gorm.DB) error {

		spesifikasi := produkPb.Spesifikasi{
			IdSpesifikasi: 0,
			Procesor: produkData.GetSpesifikasi().GetProcesor(),
			Ram: produkData.GetSpesifikasi().GetRam(),
			Penyimpanan: produkData.GetSpesifikasi().GetPenyimpanan(),
			Layar: produkData.GetSpesifikasi().GetLayar(),
			Grafis: produkData.GetSpesifikasi().GetGrafis(),
		}

		if err := tx.Table("spesifikasi").
			Where("LCASE(procesor) = ?", spesifikasi.GetProcesor()).
			FirstOrCreate(&spesifikasi).Error; err != nil{
				return err
			}

		produk := struct {
			Id uint64
			NamaProduk string
			Kategori string
			Harga float64
			Stok int32
			Deskripsi string
			Spesifikasi_id uint32
			Merek string
			TanggalRilis string
			Rating int32
		}{
			Id: produkData.GetId(),
			NamaProduk: produkData.GetNamaProduk(),
			Kategori: produkData.GetKategori(),
			Harga: produkData.GetHarga(),
			Stok: int32(produkData.GetStok()),
			Spesifikasi_id: uint32(spesifikasi.GetIdSpesifikasi()),
			Merek: produkData.GetMerek(),
			TanggalRilis: produkData.GetTanggalRilis(),
			Rating: int32(produkData.GetRating()),
		}

		if err := tx.Table("produk").Create(&produk).Error; err != nil{
			return err
		}

		Response.Id = produk.Id
		return nil
	})

	if err != nil{
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &Response, nil
}

func (p *ProdukService) UpdateProduk(ctx context.Context, produkData *produkPb.Product) (*produkPb.Status, error) {
	var Response produkPb.Status

	err := p.DB.Transaction(func(tx *gorm.DB) error {

		spesifikasi := produkPb.Spesifikasi{
			IdSpesifikasi: 0,
			Procesor: produkData.GetSpesifikasi().GetProcesor(),
			Ram: produkData.GetSpesifikasi().GetRam(),
			Penyimpanan: produkData.GetSpesifikasi().GetPenyimpanan(),
			Layar: produkData.GetSpesifikasi().GetLayar(),
			Grafis: produkData.GetSpesifikasi().GetGrafis(),
		}

		if err := tx.Table("spesifikasi").
			Where("LCASE(procesor) = ?", spesifikasi.GetProcesor()).
			FirstOrCreate(&spesifikasi).Error; err != nil{
				return err
			}

		produk := struct {
			Id uint64
			NamaProduk string
			Kategori string
			Harga float64
			Stok int32
			Deskripsi string
			Spesifikasi_id uint32
			Merek string
			TanggalRilis string
			Rating int32
		}{
			Id: produkData.GetId(),
			NamaProduk: produkData.GetNamaProduk(),
			Kategori: produkData.GetKategori(),
			Harga: produkData.GetHarga(),
			Stok: int32(produkData.GetStok()),
			Spesifikasi_id: uint32(spesifikasi.GetIdSpesifikasi()),
			Merek: produkData.GetMerek(),
			TanggalRilis: produkData.GetTanggalRilis(),
			Rating: int32(produkData.GetRating()),
		}

		if err := tx.Table("produk").Where("id = ?", produk.Id).Updates(&produk).Error; err != nil{
			return err
		}

		Response.Status = 1
		return nil
	})

	if err != nil{
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &Response, nil
}

func (p *ProdukService) DeleteProduk(ctx context.Context, id *produkPb.Product) (*produkPb.Status, error) {
	var response produkPb.Status

	if err := p.DB.Table("produk").Where("id = ?", id.GetId()).Delete(nil).Error; err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	response.Status = 1

	return &response, nil
}