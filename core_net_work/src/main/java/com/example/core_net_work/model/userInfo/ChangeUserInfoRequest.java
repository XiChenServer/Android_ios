package com.example.core_net_work.model.userInfo;

/**
 * @Author winiymissl
 * @Date 2024-02-20 18:57
 * @Version 1.0
 */
public class ChangeUserInfoRequest {
    String name;
    String nickname;
    //    String password;
    String wechat_number;
    String email;

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
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

    public ChangeUserInfoRequest(String name, String nickname, String wechat_number, String email) {
        this.name = name;
        this.nickname = nickname;
        this.wechat_number = wechat_number;
        this.email = email;
    }

    public String getWechat_number() {
        return wechat_number;
    }

    public void setWechat_number(String wechat_number) {
        this.wechat_number = wechat_number;
    }
}
