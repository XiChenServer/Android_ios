package com.example.mine_pager.adapter.model;


/**
 * @Author winiymissl
 * @Date 2024-02-12 14:16
 * @Version 1.0
 */
public class Goods {
    int image;
    String title;
    String subTitle;
    float rating;
    float price;

    public Goods(int image, String title, String subTitle, float rating, float price) {
        this.image = image;
        this.title = title;
        this.subTitle = subTitle;
        this.rating = rating;
        this.price = price;
    }

    public float getPrice() {
        return price;
    }

    public void setPrice(float price) {
        this.price = price;
    }

    public int getImage() {
        return image;
    }

    public void setImage(int image) {
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

    public float getRating() {
        return rating;
    }

    public void setRating(float rating) {
        this.rating = rating;
    }
}
