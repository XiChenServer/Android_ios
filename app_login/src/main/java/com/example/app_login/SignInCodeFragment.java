package com.example.app_login;

import android.content.Intent;
import android.os.Bundle;
import android.os.CountDownTimer;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Toast;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.fragment.app.FragmentManager;
import androidx.fragment.app.FragmentTransaction;

import com.alibaba.android.arouter.launcher.ARouter;
import com.example.app_login.databinding.FragmentSignInCodeBinding;
import com.example.core_net_work.MyRetrofit;
import com.example.core_net_work.model.login.CodeRequest;
import com.example.core_net_work.model.login.CodeResult;
import com.example.core_net_work.model.login.LoginCodeRequest;
import com.example.core_net_work.model.login.LoginRequest;
import com.example.core_net_work.model.login.LoginResult;
import com.tencent.mmkv.MMKV;

import java.util.List;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

/**
 * @Author winiymissl
 * @Date 2023-12-12 15:04
 * @Version 1.0
 */
public class SignInCodeFragment extends Fragment {
    private FragmentSignInCodeBinding binding;
    private String name;

    public SignInCodeFragment(String name) {
        this.name = name;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        //将这个登录设置为验证码登录
        View view = inflater.inflate(R.layout.fragment_sign_in_code, container, false);
        binding = FragmentSignInCodeBinding.bind(view);
        binding.btnCodeCode.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                String phone = binding.etPhoneLoginCode.getText().toString();
                if (!phone.equals("")) {
                    binding.btnCodeCode.setEnabled(false);
                    binding.btnCodeCode.setVisibility(View.GONE);
                    binding.progressBarCode.setVisibility(View.VISIBLE);
                    MyRetrofit.serviceAPI.getCode(new CodeRequest(phone)).enqueue(new Callback<CodeResult>() {
                        CountDownTimer countDownTimer;

                        @Override
                        public void onResponse(Call<CodeResult> call, Response<CodeResult> response) {
                            countDownTimer = new CountDownTimer(60000, 1000) {
                                @Override
                                public void onTick(long millisUntilFinished) {
                                    // 更新按钮上显示的文本，显示剩余时间
                                    binding.btnCodeCode.setText("(" + millisUntilFinished / 1000 + "秒)");
                                }

                                @Override
                                public void onFinish() {
                                    // 倒计时结束，启用按钮并重置文本
                                    binding.btnCodeCode.setEnabled(true);
                                    binding.btnCodeCode.setText("验证码");
                                }
                            }.start();
                            binding.btnCodeCode.setVisibility(View.VISIBLE);
                            binding.progressBarCode.setVisibility(View.GONE);
                            if (response.isSuccessful()) {
                                Toast.makeText(getActivity(), "发送成功", Toast.LENGTH_SHORT).show();
                            }
                            Toast.makeText(getActivity(), response.body().getMsg(), Toast.LENGTH_SHORT).show();
                        }

                        @Override
                        public void onFailure(Call<CodeResult> call, Throwable t) {
                            binding.btnCodeCode.setVisibility(View.VISIBLE);
                            binding.progressBarCode.setVisibility(View.GONE);
                            binding.btnCodeCode.setEnabled(true);
                            Toast.makeText(getActivity(), "网络有问题了呢", Toast.LENGTH_SHORT).show();
                        }
                    });
                } else {
                    Toast.makeText(getActivity(), "请填写手机号码", Toast.LENGTH_SHORT).show();
                }
            }
        });

        binding.btnLoginCode.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {

                String code = binding.etCodeLogin.getText().toString();
                String phone = binding.etPhoneLoginCode.getText().toString();
                if (!code.equals("") && !phone.equals("")) {
                    binding.btnLoginCode.setVisibility(View.GONE);
                    binding.progressCodeBarLogin.setVisibility(View.VISIBLE);
                    MyRetrofit.serviceAPI.login_code(new LoginCodeRequest(code, phone)).enqueue(new Callback<LoginResult>() {


                        @Override
                        public void onResponse(Call<LoginResult> call, Response<LoginResult> response) {
                            binding.progressCodeBarLogin.setVisibility(View.GONE);
                            binding.btnLoginCode.setVisibility(View.VISIBLE);
                            if (response.isSuccessful()) {
                                //本地存储token
                                MMKV mmkv = MMKV.defaultMMKV();
                                mmkv.encode("token", response.body().getData().getToken());
                                Toast.makeText(getActivity(), "登录成功", Toast.LENGTH_SHORT).show();
                                ARouter.getInstance().build("/sellcowhourse/app_MainActivity").withFlags(Intent.FLAG_ACTIVITY_NEW_TASK | Intent.FLAG_ACTIVITY_CLEAR_TASK).navigation();
                                getActivity().overridePendingTransition(com.example.common.R.anim.set_in, com.example.common.R.anim.set_out);
                            }
                            Toast.makeText(getActivity(), "出问题了呢", Toast.LENGTH_SHORT).show();
                        }

                        @Override
                        public void onFailure(Call<LoginResult> call, Throwable t) {
                            binding.progressCodeBarLogin.setVisibility(View.GONE);
                            binding.btnLoginCode.setVisibility(View.VISIBLE);
                            Toast.makeText(getActivity(), "网络有问题了呢", Toast.LENGTH_SHORT).show();
                        }
                    });
                } else {
                    Toast.makeText(getActivity(), "请完善正确的信息", Toast.LENGTH_SHORT).show();
                }
            }
        });
        binding.chipCode.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                FragmentManager supportFragmentManager = getActivity().getSupportFragmentManager();
                SignInPsswrdFragment signInPsswrdFragment = new SignInPsswrdFragment("登录");
                FragmentTransaction fragmentTransaction = supportFragmentManager.beginTransaction();
                fragmentTransaction.setCustomAnimations(android.R.anim.fade_in, android.R.anim.fade_out);
                fragmentTransaction.replace(R.id.frame_in_code_psswrd, signInPsswrdFragment);
                fragmentTransaction.commit();
                binding.frameInCodeCode.setVisibility(View.GONE);
            }
        });
        return view;
    }
}
