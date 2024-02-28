package com.example.core_net_work;

import com.example.core_net_work.model.goods.BoughtResult;
import com.example.core_net_work.model.goods.ProductCollectCancelResult;
import com.example.core_net_work.model.goods.ProductCollectResult;
import com.example.core_net_work.model.goods.ProductLikeCancelResul;
import com.example.core_net_work.model.goods.ProductLikeResult;
import com.example.core_net_work.model.goods.ProductSimpleInfoResult;
import com.example.core_net_work.model.goods.SearchResult;
import com.example.core_net_work.model.goods.SoldInfoResult;
import com.example.core_net_work.model.goods.UserLoadedResult;
import com.example.core_net_work.model.login.CodeRequest;
import com.example.core_net_work.model.login.CodeResult;
import com.example.core_net_work.model.login.LoginCodeRequest;
import com.example.core_net_work.model.login.LoginRequest;
import com.example.core_net_work.model.login.LoginResult;
import com.example.core_net_work.model.login.RegisterRequest;
import com.example.core_net_work.model.login.RegisterResult;
import com.example.core_net_work.model.userInfo.ChangePhoneResult;
import com.example.core_net_work.model.userInfo.ChangeUserInfoRequest;
import com.example.core_net_work.model.userInfo.ChangeUserInfoResult;
import com.example.core_net_work.model.userInfo.UploadAddressResult;
import com.example.core_net_work.model.userInfo.UploadAvatarResult;
import com.example.core_net_work.model.userInfo.UserInfoResult;

import java.util.List;

import okhttp3.MultipartBody;
import okhttp3.RequestBody;
import retrofit2.Call;
import retrofit2.http.Body;
import retrofit2.http.Field;
import retrofit2.http.FormUrlEncoded;
import retrofit2.http.GET;
import retrofit2.http.Header;
import retrofit2.http.Multipart;
import retrofit2.http.POST;
import retrofit2.http.Part;
import retrofit2.http.Query;

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

    @FormUrlEncoded
    @POST("/user/changes/mobile/phone")
    Call<ChangePhoneResult> changePhoneNumber(@Header("Authorization") String token, @Field("phone_number") String phoneNumber, @Field("verification_code") String verificationCode);

    @POST("/user/modify/info")
    Call<ChangeUserInfoResult> changeUserInfo(@Header("Authorization") String token, @Body ChangeUserInfoRequest user);

    @POST("/user/collect/product")
    Call<ProductCollectResult> collectProduct(@Header("Authorization") String token, @Query("commodity_identity ") String commodity_identity);

    @POST("/user/uncollect/product")
    Call<ProductCollectCancelResult> cancelCollectProduct(@Header("Authorization") String token, @Query("commodity_identity") String commodity_identity);

    @POST("/user/like/product")
    Call<ProductLikeResult> likeProduct(@Header("Authorization") String token, @Query("commodity_identity") String commodity_identity);

    @POST("/user/unlike/product")
    Call<ProductLikeCancelResul> cancelLikeProduct(@Header("Authorization") String token, @Query("commodity_identity_") String commodity_identity_);

    @POST("/user/upload/address")
    Call<UploadAddressResult> uploadAddress(@Header("Authorization") String token, @Body String country, @Body String province, @Body String city, @Body String street, @Body String contact, @Body String post_code, @Body String identity);

    @POST("search/product")
    Call<SearchResult> searchProduct(@Query("data") String data);

    @Multipart
    @POST("/user/uploads/avatar")
    Call<UploadAvatarResult> uploadAvatar(@Header("Authorization") String token, @Part MultipartBody.Part file);

    @POST("/user/adds/products")
    @Multipart
    Call<ProductSimpleInfoResult> addProduct(@Header("Authorization") String token,
                                             @Part("type") List<RequestBody> type,
                                             @Part("title") RequestBody title,
                                             @Part("number") RequestBody number,
                                             @Part("information") RequestBody information,
                                             @Part("price") RequestBody price,
                                             @Part("is_auction") RequestBody is_auction,
                                             @Part("country") RequestBody country,
                                             @Part("province") RequestBody province,
                                             @Part("city") RequestBody city,
                                             @Part("contact") RequestBody contact,
                                             @Part("post_code") RequestBody post_code,
                                             @Part List<MultipartBody.Part> files);

    @GET("/products/simple_info")
    Call<ProductSimpleInfoResult> getProductSimpleInfo();

    @POST("/get/user_all_pro_list")
    Call<UserLoadedResult> getUserUploaded(@Query("user_identity") String user_identity);

    @POST("/user/order/find/AllSellOrders")
    Call<SoldInfoResult> getSoldInfo(@Header("Authorization") String token);

    @POST("/user/order/find/AllBuyOrder")
    Call<BoughtResult> getBoughtInfo(@Header("Authorization") String token);

}

