package com.example.core_net_work.model.login;

import com.example.core_net_work.model.BaseResult;

/**
 * @Author winiymissl
 * @Date 2023-12-11 13:35
 * @Version 1.0
 */
public class RegisterResult extends BaseResult {
    RegisterDataResult data;

    public RegisterDataResult getData() {
        return data;
    }

    public void setData(RegisterDataResult data) {
        this.data = data;
    }

    class RegisterDataResult {
        String nickname;
        String account;

        public String getNickname() {
            return nickname;
        }

        public void setNickname(String nickname) {
            this.nickname = nickname;
        }

        public String getAccount() {
            return account;
        }

        public void setAccount(String account) {
            this.account = account;
        }
    }
}

