package com.example.common.room.dao;

import androidx.room.Dao;
import androidx.room.Delete;
import androidx.room.Insert;
import androidx.room.Query;

import com.example.common.room.entitues.ShoppingCarOrder;

import java.util.List;

/**
 * @Author winiymissl
 * @Date 2024-02-22 18:16
 * @Version 1.0
 */
@Dao
public interface ShoppingCarDao {
    @Query("SELECT * FROM shopping_car_order")
    List<ShoppingCarOrder> getAllInfo();

    @Insert()
    void insertAll(List<ShoppingCarOrder> shoppingCarOrders);

    @Delete()
    void delete(ShoppingCarOrder shoppingCarOrder);

}

