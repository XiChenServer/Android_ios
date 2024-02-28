package com.example.common.room;

/**
 * @Author winiymissl
 * @Date 2024-02-23 20:51
 * @Version 1.0
 */
public class UserAddress {
    String country;
    String province;
    String city;

    @Override
    public String toString() {
        return "UserAddress{" +
                "country='" + country + '\'' +
                ", province='" + province + '\'' +
                ", city='" + city + '\'' +
                ", street='" + street + '\'' +
                ", identity='" + identity + '\'' +
                ", post_code='" + post_code + '\'' +
                ", contact='" + contact + '\'' +
                '}';
    }

    String street;
    String identity;
    String post_code;
    String contact;


    public String getCity() {
        return city;
    }

    public void setCity(String city) {
        this.city = city;
    }

    public String getContact() {
        return contact;
    }

    public void setContact(String contact) {
        this.contact = contact;
    }

    public String getCountry() {
        return country;
    }

    public void setCountry(String country) {
        this.country = country;
    }

    public String getIdentity() {
        return identity;
    }

    public void setIdentity(String identity) {
        this.identity = identity;
    }

    public String getPost_code() {
        return post_code;
    }

    public void setPost_code(String post_code) {
        this.post_code = post_code;
    }

    public String getStreet() {
        return street;
    }

    public void setStreet(String street) {
        this.street = street;
    }

    public String getProvince() {
        return province;
    }

    public void setProvince(String province) {
        this.province = province;
    }
}
