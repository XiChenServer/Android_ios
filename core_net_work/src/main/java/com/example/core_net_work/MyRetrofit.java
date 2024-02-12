package com.example.core_net_work;

import retrofit2.Retrofit;
import retrofit2.converter.gson.GsonConverterFactory;

/**
 * @Author winiymissl
 * @Date 2023-12-11 13:14
 * @Version 1.0
 */
public class MyRetrofit {

    public static final Retrofit retrofit = new Retrofit.Builder().baseUrl("http://8.130.86.26:13000/").addConverterFactory(GsonConverterFactory.create()).build();
    public static final ServiceAPI serviceAPI = retrofit.create(ServiceAPI.class);
}
