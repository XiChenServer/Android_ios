package com.example.core_net_work.model.login;

/**
 * @Author winiymissl
 * @Date 2023-12-12 19:43
 * @Version 1.0
 */
public class LoginCodeRequest {
    String verification_code;
    String phone_number;

    public String getVerification_code() {
        return verification_code;
    }

    public void setVerification_code(String verification_code) {
        this.verification_code = verification_code;
    }

    public String getPhone_number() {
        return phone_number;
    }

    public void setPhone_number(String phone_number) {
        this.phone_number = phone_number;
    }

    public LoginCodeRequest(String verification_code, String phone_number) {
        this.verification_code = verification_code;
        this.phone_number = phone_number;
    }
}
