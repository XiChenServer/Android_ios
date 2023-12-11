package com.example.core_net_work.model.login;

/**
 * @Author winiymissl
 * @Date 2023-12-11 13:37
 * @Version 1.0
 */
public class CodeRequest {
    String phone;

    public String getPhone() {
        return phone;
    }

    public void setPhone(String phone) {
        this.phone = phone;
    }

    public CodeRequest(String phone) {
        this.phone = phone;
    }
}
