package com.example.common.room.dao;

import androidx.room.Dao;
import androidx.room.Delete;
import androidx.room.Insert;
import androidx.room.Query;

import com.example.common.room.entitues.ProductSimple;

import java.util.List;

/**
 * @Author winiymissl
 * @Date 2024-02-21 17:19
 * @Version 1.0
 */
@Dao
public interface ProductSimpleDao {
    @Query("SELECT * FROM product_simple")
    List<ProductSimple> getAllInfo();

    @Insert()
    void insertAll(List<ProductSimple> productSimples);

    @Delete()
    void delete(ProductSimple productSimple);

    @Query("DELETE FROM product_simple")
    void deleteAll();
}
