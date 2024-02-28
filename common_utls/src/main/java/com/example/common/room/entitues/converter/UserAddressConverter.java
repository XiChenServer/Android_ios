package com.example.common.room.entitues.converter;

import androidx.room.TypeConverter;

import com.example.common.room.UserAddress;
import com.google.gson.Gson;

/**
 * @Author winiymissl
 * @Date 2024-02-23 21:00
 * @Version 1.0
 */
public class UserAddressConverter {
    private static final Gson gson = new Gson();

    @TypeConverter
    public static UserAddress fromString(String value) {
        return value == null ? null : gson.fromJson(value, UserAddress.class);
    }

    @TypeConverter
    public static String addressToString(UserAddress address) {
        return address == null ? null : gson.toJson(address);
    }
}
