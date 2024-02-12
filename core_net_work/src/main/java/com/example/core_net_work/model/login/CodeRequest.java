package com.example.core_net_work.model.login;

/**
 * @Author winiymissl
 * @Date 2023-12-11 13:37
 * @Version 1.0
 */
public class CodeRequest {
    String phone_number;

    public String getPhone_number() {
        return phone_number;
    }

    public void setPhone_number(String phone_number) {
        this.phone_number = phone_number;
    }

    public CodeRequest(String phone) {
        this.phone_number = phone;
    }
}
