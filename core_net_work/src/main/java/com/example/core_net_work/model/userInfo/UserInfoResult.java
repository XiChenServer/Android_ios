package com.example.core_net_work.model.userInfo;

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
        String avatar;

        public String getAvatar() {
            return avatar;
        }

        public void setAvatar(String avatar) {
            this.avatar = avatar;
        }
    }
}
