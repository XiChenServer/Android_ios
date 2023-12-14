package com.example.core_net_work.model.login;

import com.example.core_net_work.model.BaseResult;

/**
 * @Author winiymissl
 * @Date 2023-12-12 16:18
 * @Version 1.0
 */
public class LoginResult extends BaseResult {
    LoginResponseData data;

    public LoginResponseData getData() {
        return data;
    }

    @Override
    public String toString() {
        return "LoginResult{" + "data=" + data + '}';
    }

    public void setData(LoginResponseData data) {
        this.data = data;
    }

    public class LoginResponseData {
        @Override
        public String toString() {
            return "LoginResponseData{" + "token='" + token + '\'' + '}';
        }

        String token;

        public String getToken() {
            return token;
        }

        public void setToken(String token) {
            this.token = token;
        }
    }
}
