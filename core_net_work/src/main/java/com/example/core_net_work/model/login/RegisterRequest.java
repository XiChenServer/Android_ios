package com.example.core_net_work.model.login;

/**
 * @Author winiymissl
 * @Date 2023-12-11 13:24
 * @Version 1.0
 */
public class RegisterRequest {
    String phone_number;
    String password;
    String verification_code;

    public RegisterRequest(String phone, String password, String phone_code) {
        this.phone_number = phone;
        this.password = password;
        this.verification_code = phone_code;
    }

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

    public String getVerification_code() {
        return verification_code;
    }

    public void setVerification_code(String verification_code) {
        this.verification_code = verification_code;
    }
}
