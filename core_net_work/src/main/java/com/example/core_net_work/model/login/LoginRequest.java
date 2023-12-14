package com.example.core_net_work.model.login;

/**
 * @Author winiymissl
 * @Date 2023-12-12 16:18
 * @Version 1.0
 */
public class LoginRequest {
    String phone_number;

    String password;

    public String getPhone_number() {
        return phone_number;
    }

    public void setPhone_number(String phone_number) {
        this.phone_number = phone_number;
    }

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }

    public LoginRequest(String phone_number, String password) {
        this.phone_number = phone_number;
        this.password = password;
    }
}
