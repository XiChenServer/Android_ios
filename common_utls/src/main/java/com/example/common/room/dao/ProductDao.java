package com.example.common.room.dao;

import androidx.room.Dao;
import androidx.room.Delete;
import androidx.room.Insert;
import androidx.room.Query;

import com.example.common.room.entitues.Product;

import java.util.List;

/**
 * @Author winiymissl
 * @Date 2024-02-21 17:36
 * @Version 1.0
 */
@Dao
public interface ProductDao {
    @Query("SELECT * FROM product")
    List<Product> getAllInfo();

    @Insert()
    void insertAll(List<Product> products);

    @Delete()
    void delete(Product product);

    @Query("SELECT * FROM product WHERE commodity_identity = :identity")
    List<Product> getProduct(String identity);
}
