package com.example.common.room.dao;

import androidx.room.Dao;
import androidx.room.Delete;
import androidx.room.Insert;
import androidx.room.Query;

import com.example.common.room.entitues.User;

import java.util.List;

/**
 * @Author winiymissl
 * @Date 2024-02-12 22:21
 * @Version 1.0
 */
@Dao
public interface UserDao {
    @Query("SELECT * FROM User")
    List<User> getAllInfo();

    @Insert()
    void insertAll(User... users);

    @Delete()
    void delete(List<User> user);

    @Query("UPDATE user SET name = :newName")
    void changeUserInfo(String newName);
}
