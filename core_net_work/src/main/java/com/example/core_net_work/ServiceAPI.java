package com.example.core_net_work;

import com.example.core_net_work.model.login.CodeRequest;
import com.example.core_net_work.model.login.CodeResult;
import com.example.core_net_work.model.login.LoginCodeRequest;
import com.example.core_net_work.model.login.LoginRequest;
import com.example.core_net_work.model.login.LoginResult;
import com.example.core_net_work.model.login.RegisterRequest;
import com.example.core_net_work.model.login.RegisterResult;
import com.example.core_net_work.model.userInfo.UserInfoResult;

import retrofit2.Call;
import retrofit2.http.Body;
import retrofit2.http.GET;
import retrofit2.http.Header;
import retrofit2.http.POST;

/**
 * @Author winiymissl
 * @Date 2023-12-11 13:14
 * @Version 1.0
 */
public interface ServiceAPI {
    @POST("/send_phone_code")
    Call<CodeResult> getCode(@Body CodeRequest code);

    @POST("/user/register/phone")
    Call<RegisterResult> register(@Body RegisterRequest myRegister);

    @POST("/user/login/phone_and_password")
    Call<LoginResult> login_psswrd(@Body LoginRequest loginRequest);

    @POST("/user/login/phone")
    Call<LoginResult> login_code(@Body LoginCodeRequest loginCodeRequest);

    @GET("/user/get/info")
    Call<UserInfoResult> getUserInfo(@Header("Authorization") String token);


}
