package com.example.common.room;

import android.content.Context;

import androidx.room.Database;
import androidx.room.Room;
import androidx.room.RoomDatabase;
import androidx.room.TypeConverters;

import com.example.common.room.dao.ProductDao;
import com.example.common.room.dao.ProductSimpleDao;
import com.example.common.room.dao.ShoppingCarDao;
import com.example.common.room.dao.UserDao;
import com.example.common.room.entitues.Product;
import com.example.common.room.entitues.ProductSimple;
import com.example.common.room.entitues.ShoppingCarOrder;
import com.example.common.room.entitues.User;
import com.example.common.room.entitues.converter.UserAddressConverter;

/**
 * @Author winiymissl
 * @Date 2024-02-19 16:18
 * @Version 1.0
 */
@Database(entities = {User.class, ProductSimple.class, Product.class, ShoppingCarOrder.class}, version = 19, exportSchema = false)
@TypeConverters({UserAddressConverter.class})

public abstract class AppDatabase extends RoomDatabase {
    public abstract UserDao userDao();

    public abstract ProductSimpleDao productSimpleDao();

    public abstract ShoppingCarDao shoppingCarDao();

    public abstract ProductDao productDao();

    private static AppDatabase instance;

    public static synchronized AppDatabase getInstance(Context context) {
        if (instance == null) {
            instance = Room.databaseBuilder(context.getApplicationContext(),
                            AppDatabase.class, "app_database")
                    .fallbackToDestructiveMigration()
                    .build();
        }
        return instance;
    }
}
