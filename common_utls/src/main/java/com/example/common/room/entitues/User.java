package com.example.common.room.entitues;

import androidx.room.ColumnInfo;
import androidx.room.Entity;
import androidx.room.PrimaryKey;
import androidx.room.TypeConverters;

import com.example.common.room.UserAddress;
import com.example.common.room.entitues.converter.UserAddressConverter;

/**
 * @Author winiymissl
 * @Date 2024-02-19 15:58
 * @Version 1.0
 */
@Entity(tableName = "user")
@TypeConverters(UserAddressConverter.class)

public class User {
    @PrimaryKey(autoGenerate = true)
    public int uid;

    @ColumnInfo(name = "nickname")
    public String nickname;

    @ColumnInfo(name = "avatar", typeAffinity = ColumnInfo.TEXT)
    public String avatar;
    @ColumnInfo(name = "email", typeAffinity = ColumnInfo.TEXT)
    public String email;
    @ColumnInfo(name = "phone_number", typeAffinity = ColumnInfo.TEXT)
    public String phone_number;
    @ColumnInfo(name = "name", typeAffinity = ColumnInfo.TEXT)
    public String name;
    @ColumnInfo(name = "wechat_number", typeAffinity = ColumnInfo.TEXT)
    public String wechat_number;

    public String getWechat_number() {
        return wechat_number;
    }

    public void setWechat_number(String wechat_number) {
        this.wechat_number = wechat_number;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String getUser_identity() {
        return user_identity;
    }

    public void setUser_identity(String user_identity) {
        this.user_identity = user_identity;
    }

    @ColumnInfo(name = "user_identity", typeAffinity = ColumnInfo.TEXT)
    public String user_identity;
    @ColumnInfo(name = "account", typeAffinity = ColumnInfo.TEXT)
    public String account;
    @ColumnInfo(name = "address", typeAffinity = ColumnInfo.TEXT)
    public UserAddress address;

    public UserAddress getAddress() {
        return address;
    }

    public void setAddress(UserAddress address) {
        this.address = address;
    }

    public User(String nickname, String avatar, String email, String phone_number, String name, String wechat_number, String user_identity, String account, UserAddress address) {
        this.nickname = nickname;
        this.avatar = avatar;
        this.email = email;
        this.phone_number = phone_number;
        this.name = name;
        this.wechat_number = wechat_number;
        this.user_identity = user_identity;
        this.account = account;
        this.address = address;
    }

    public int getUid() {
        return uid;
    }

    public void setUid(int uid) {
        this.uid = uid;
    }

    public String getNickname() {
        return nickname;
    }

    public void setNickname(String nickname) {
        this.nickname = nickname;
    }

    public String getAvatar() {
        return avatar;
    }

    public void setAvatar(String avatar) {
        this.avatar = avatar;
    }

    public String getPhone_number() {
        return phone_number;
    }

    public void setPhone_number(String phone_number) {
        this.phone_number = phone_number;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getAccount() {
        return account;
    }

    public void setAccount(String account) {
        this.account = account;
    }
}
