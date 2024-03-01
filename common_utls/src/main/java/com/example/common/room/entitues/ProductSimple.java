package com.example.common.room.entitues;

import androidx.room.ColumnInfo;
import androidx.room.Entity;
import androidx.room.PrimaryKey;

/**
 * @Author winiymissl
 * @Date 2024-02-21 17:20
 * @Version 1.0
 */
@Entity(tableName = "product_simple")
public class ProductSimple {
    @PrimaryKey(autoGenerate = true)
    public int uid;
    //商品的唯一标识
    @ColumnInfo(name = "product_identity", typeAffinity = ColumnInfo.TEXT)
    String product_identity;
    @ColumnInfo(name = "image", typeAffinity = ColumnInfo.TEXT)
    public String image;
    @ColumnInfo(name = "title", typeAffinity = ColumnInfo.TEXT)
    public String title;
    @ColumnInfo(name = "subTitle", typeAffinity = ColumnInfo.TEXT)
    public String subTitle;
    @ColumnInfo(name = "price", typeAffinity = ColumnInfo.TEXT)
    public String price;
    @ColumnInfo(name = "rating", typeAffinity = ColumnInfo.TEXT)
    public String rating;
    @ColumnInfo(name = "ID", typeAffinity = ColumnInfo.TEXT)
    public String id;
    @ColumnInfo(name = "information", typeAffinity = ColumnInfo.TEXT)
    public String information;

    public ProductSimple(String product_identity, String image, String title, String information, String price, String rating, String id) {
        this.product_identity = product_identity;
        this.image = image;
        this.title = title;
        this.information = information;
        this.price = price;
        this.rating = rating;
        this.id = id;
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getProduct_identity() {
        return product_identity;
    }

    public void setProduct_identity(String product_identity) {
        this.product_identity = product_identity;
    }

    public String getImage() {
        return image;
    }

    public void setImage(String image) {
        this.image = image;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public String getSubTitle() {
        return subTitle;
    }

    public void setSubTitle(String subTitle) {
        this.subTitle = subTitle;
    }

    public String getPrice() {
        return price;
    }

    public void setPrice(String price) {
        this.price = price;
    }

    public String getRating() {
        return rating;
    }

    public void setRating(String rating) {
        this.rating = rating;
    }
}
