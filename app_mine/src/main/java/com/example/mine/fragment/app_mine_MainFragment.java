package com.example.mine.fragment;

import android.content.Intent;
import android.graphics.drawable.Drawable;
import android.os.Bundle;
import android.util.Log;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Toast;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;

import com.alibaba.android.arouter.facade.annotation.Route;
import com.alibaba.android.arouter.launcher.ARouter;
import com.baoyz.actionsheet.ActionSheet;
import com.bumptech.glide.Glide;
import com.bumptech.glide.load.DataSource;
import com.bumptech.glide.load.engine.GlideException;
import com.bumptech.glide.request.RequestListener;
import com.bumptech.glide.request.target.Target;
import com.example.core_net_work.MyRetrofit;
import com.example.core_net_work.model.userInfo.UserInfoResult;
import com.example.mine.R;
import com.example.mine.adpater.MineListViewAdapter;
import com.example.mine.adpater.impl.MineListViewAdapterImpl;
import com.example.mine.databinding.AppMineFragmentBinding;
import com.google.android.material.snackbar.Snackbar;
import com.tencent.mmkv.MMKV;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

/**
 * @Author winiymissl
 * @Date 2023-12-12 17:54
 * @Version 1.0
 */
@Route(path = "/fragment/app_mine_MainFragment")
public class app_mine_MainFragment extends Fragment {
    AppMineFragmentBinding binding;

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.app_mine_fragment, container, false);
        binding = AppMineFragmentBinding.bind(view);
        MineListViewAdapter adapter = new MineListViewAdapter(MineListViewAdapterImpl.getList());
        binding.mineListview.setAdapter(adapter);

        binding.ivMine.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                ActionSheet.createBuilder(getActivity(), getActivity().getSupportFragmentManager()).setCancelButtonTitle("Cancel").setOtherButtonTitles("设置背景图片").setCancelableOnTouchOutside(true).setListener(new ActionSheet.ActionSheetListener() {
                    @Override
                    public void onDismiss(ActionSheet actionSheet, boolean isCancel) {

                    }

                    @Override
                    public void onOtherButtonClick(ActionSheet actionSheet, int index) {
                        if (index == 1) {
                            //跳转到照片选择器
                        }
                    }
                }).show();
            }
        });
        binding.btnLogout.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                MMKV mmkv = MMKV.defaultMMKV();
                if (mmkv.getBoolean("token", false)) {
                    mmkv.remove("token");
                    binding.constraintIusLogin.setVisibility(View.GONE);
                    binding.textViewIsLogin.setVisibility(View.VISIBLE);
                }
                Snackbar.make(requireView(), "是否登录", 1000).setAction("去登陆", new View.OnClickListener() {
                    @Override
                    public void onClick(View v) {
                        ARouter.getInstance().build("/app_login/AppLoginMainActivity").withFlags(Intent.FLAG_ACTIVITY_NEW_TASK | Intent.FLAG_ACTIVITY_CLEAR_TASK).navigation();
                        getActivity().overridePendingTransition(com.example.common.R.anim.set_in, com.example.common.R.anim.set_out);
                    }
                }).show();
            }
        });
        String token = MMKV.defaultMMKV().getString("token", null);
        if (token != null) {
            Log.d("ThereIsProblem", token);
            MyRetrofit.serviceAPI.getUserInfo("bearer " + token).enqueue(new Callback<UserInfoResult>() {
                @Override
                public void onResponse(Call<UserInfoResult> call, Response<UserInfoResult> response) {
                    if (response.isSuccessful()) {
                        Toast.makeText(getActivity(), "返回成功", Toast.LENGTH_SHORT).show();
                        Log.d("ThereIsProblem", response.body().getData().getAvatar());
                        //直接使用Glide，将图像显示
                        Glide.with(app_mine_MainFragment.this).load(response.body().getData().getAvatar()).error(com.example.common.R.drawable.default_avater).listener(new RequestListener<Drawable>() {
                            @Override
                            public boolean onLoadFailed(@Nullable GlideException e, @Nullable Object model, @NonNull Target<Drawable> target, boolean isFirstResource) {
                                Log.d("ThereIsProblem", e.toString());
                                if (e instanceof GlideException) {
                                    GlideException glideException = (GlideException) e;
                                    glideException.logRootCauses("Glide Loading Error - Detailed Log");
                                }
                                return false;
                            }

                            @Override
                            public boolean onResourceReady(@NonNull Drawable resource, @NonNull Object model, Target<Drawable> target, @NonNull DataSource dataSource, boolean isFirstResource) {
                                return false;
                            }
                        }).into(binding.imageView);
                    } else {
                        Toast.makeText(getActivity(), "返回失败", Toast.LENGTH_SHORT).show();
                    }
                }

                @Override
                public void onFailure(Call<UserInfoResult> call, Throwable t) {
                    Toast.makeText(getActivity(), "返回失败" + t.toString(), Toast.LENGTH_SHORT).show();
                    Log.d("ThereIsProblem", t.toString());
                }
            });
        } else {
            Toast.makeText(getActivity(), "token为空", Toast.LENGTH_SHORT).show();
        }
        return view;
    }
}
