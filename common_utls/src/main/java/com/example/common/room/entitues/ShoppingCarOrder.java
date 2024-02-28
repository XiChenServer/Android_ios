package com.example.common.room.entitues;

import androidx.room.ColumnInfo;
import androidx.room.Entity;
import androidx.room.PrimaryKey;

/**
 * @Author winiymissl
 * @Date 2024-02-22 14:19
 * @Version 1.0
 */
@Entity(tableName = "shopping_car_order")
public class ShoppingCarOrder {
    @PrimaryKey(autoGenerate = true)
    public int uid;
    @ColumnInfo(name = "image", typeAffinity = ColumnInfo.TEXT)

    public String image;
    @ColumnInfo(name = "price", typeAffinity = ColumnInfo.REAL)
    public float price;
    @ColumnInfo(name = "count", typeAffinity = ColumnInfo.INTEGER)
    public int count;
    @ColumnInfo(name = "name", typeAffinity = ColumnInfo.TEXT)
    public String name;

    public ShoppingCarOrder(String image, float price, int count, String name) {
        this.image = image;
        this.price = price;
        this.count = count;
        this.name = name;
    }

    public String getImage() {
        return image;
    }

    public void setImage(String image) {
        this.image = image;
    }

    public float getPrice() {
        return price;
    }

    public void setPrice(float price) {
        this.price = price;
    }

    public int getCount() {
        return count;
    }

    public void setCount(int count) {
        this.count = count;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }
}
