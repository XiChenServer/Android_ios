package com.example.core_net_work.model.userInfo;

import com.example.common.room.UserAddress;
import com.example.core_net_work.model.BaseResult;

/**
 * @Author winiymissl
 * @Date 2023-12-18 19:44
 * @Version 1.0
 */
public class UserInfoResult extends BaseResult {
    UserInfoAllResult data;

    public UserInfoAllResult getData() {
        return data;
    }

    public void setData(UserInfoAllResult data) {
        this.data = data;
    }


    public class UserInfoAllResult {
        String account;

        public String getUser_identity() {
            return user_identity;
        }

        public void setUser_identity(String user_identity) {
            this.user_identity = user_identity;
        }

        String user_identity;

        public String getAccount() {
            return account;
        }

        public void setAccount(String account) {
            this.account = account;
        }

        UserAddress address;
        String name;
        String nickname;
        String phone_number;
        String avatar;
        String wechat_number;
        String email;
        public UserAddress getAddress() {
            return address;
        }

        public void setAddress(UserAddress address) {
            this.address = address;
        }

        public String getName() {
            return name;
        }

        public void setName(String name) {
            this.name = name;
        }

        public String getNickname() {
            return nickname;
        }

        public void setNickname(String nickname) {
            this.nickname = nickname;
        }

        public String getPhone_number() {
            return phone_number;
        }

        public void setPhone_number(String phone_number) {
            this.phone_number = phone_number;
        }

        public String getWechat_number() {
            return wechat_number;
        }

        public void setWechat_number(String wechat_number) {
            this.wechat_number = wechat_number;
        }

        public String getEmail() {
            return email;
        }

        public void setEmail(String email) {
            this.email = email;
        }

        public String getAvatar() {
            return avatar;
        }

        public void setAvatar(String avatar) {
            this.avatar = avatar;
        }
    }
}
