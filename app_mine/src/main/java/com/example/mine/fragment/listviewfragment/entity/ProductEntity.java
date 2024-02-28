package com.example.mine.fragment.listviewfragment.entity;

/**
 * @Author winiymissl
 * @Date 2024-02-26 21:11
 * @Version 1.0
 */
public class ProductEntity {
    String image;
    String name;
    String price;

    public ProductEntity(String image, String name, String price) {
        this.image = image;
        this.name = name;
        this.price = price;
    }

    public String getImage() {
        return image;
    }

    public void setImage(String image) {
        this.image = image;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getPrice() {
        return price;
    }

    public void setPrice(String price) {
        this.price = price;
    }
}
