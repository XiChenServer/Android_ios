package com.example.common.room.entitues;

import androidx.room.ColumnInfo;
import androidx.room.Entity;
import androidx.room.PrimaryKey;

/**
 * @Author winiymissl
 * @Date 2024-02-21 17:36
 * @Version 1.0
 */
@Entity(tableName = "product")
public class Product {
    @PrimaryKey(autoGenerate = true)
    public int uid;

    //商品的唯一标识
    @ColumnInfo(name = "commodity_identity", typeAffinity = ColumnInfo.TEXT)
    public String commodity_identity;

    @ColumnInfo(name = "ID", typeAffinity = ColumnInfo.TEXT)
    public String ID;

    public String getProduct_identity() {
        return commodity_identity;
    }

    public void setProduct_identity(String product_identity) {
        this.commodity_identity = product_identity;
    }

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
    @ColumnInfo(name = "is_auction", typeAffinity = ColumnInfo.TEXT)
    public String is_auction;//是否拍卖
    @ColumnInfo(name = "number", typeAffinity = ColumnInfo.TEXT)
    public String number;

    public String getCommodity_identity() {
        return commodity_identity;
    }

    public void setCommodity_identity(String commodity_identity) {
        this.commodity_identity = commodity_identity;
    }

    public String getNumber() {
        return number;
    }

    public void setNumber(String number) {
        this.number = number;
    }

    @ColumnInfo(name = "type", typeAffinity = ColumnInfo.TEXT)
    public String type;

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }

    @ColumnInfo(name = "street", typeAffinity = ColumnInfo.TEXT)
    public String street;
    @ColumnInfo(name = "city", typeAffinity = ColumnInfo.TEXT)
    public String city;
    @ColumnInfo(name = "country", typeAffinity = ColumnInfo.TEXT)
    public String country;
    @ColumnInfo(name = "province", typeAffinity = ColumnInfo.TEXT)
    public String province;
    @ColumnInfo(name = "contact", typeAffinity = ColumnInfo.TEXT)
    public String contact;
    @ColumnInfo(name = "post_code", typeAffinity = ColumnInfo.TEXT)
    public String post_code;
    @ColumnInfo(name = "file", typeAffinity = ColumnInfo.TEXT)
    public String file;

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

    public String getIs_auction() {
        return is_auction;
    }

    public void setIs_auction(String is_auction) {
        this.is_auction = is_auction;
    }

    public String getStreet() {
        return street;
    }

    public void setStreet(String street) {
        this.street = street;
    }

    public String getCity() {
        return city;
    }

    public void setCity(String city) {
        this.city = city;
    }

    public String getCountry() {
        return country;
    }

    public void setCountry(String country) {
        this.country = country;
    }

    public String getProvince() {
        return province;
    }

    public void setProvince(String province) {
        this.province = province;
    }

    public String getContact() {
        return contact;
    }

    public void setContact(String contact) {
        this.contact = contact;
    }

    public String getPost_code() {
        return post_code;
    }

    public void setPost_code(String post_code) {
        this.post_code = post_code;
    }

    public String getFile() {
        return file;
    }

    public void setFile(String file) {
        this.file = file;
    }
}
