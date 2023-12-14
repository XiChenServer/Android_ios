package com.example.app_login;

import android.content.Intent;
import android.os.Bundle;

import androidx.fragment.app.Fragment;
import androidx.fragment.app.FragmentManager;
import androidx.fragment.app.FragmentTransaction;

import android.util.Log;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Toast;

import com.alibaba.android.arouter.launcher.ARouter;
import com.example.app_login.databinding.FragmentSignInPasswrdBinding;
import com.example.core_net_work.MyRetrofit;
import com.example.core_net_work.model.login.LoginRequest;
import com.example.core_net_work.model.login.LoginResult;
import com.tencent.mmkv.MMKV;

import java.util.List;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;


public class SignInPsswrdFragment extends Fragment {
    private String name;
    FragmentSignInPasswrdBinding binding;

    public SignInPsswrdFragment(String name) {
        this.name = name;
        // Required empty public constructor
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public SignInPsswrdFragment() {

    }

    @Override
    public View onCreateView(LayoutInflater inflater, ViewGroup container, Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_sign_in_passwrd, container, false);
        binding = FragmentSignInPasswrdBinding.bind(view);
        binding.btnLoginPsswrd.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                String phone = binding.etPhoneNumLogin.getText().toString();
                String password = binding.etPsswrdLogin.getText().toString();
                Log.d("IsThereProblem", phone + password);
                if (!phone.equals("") && !password.equals("")) {
                    binding.btnLoginPsswrd.setVisibility(View.GONE);
                    binding.progressRegisterBarLogin.setVisibility(View.VISIBLE);
                    MyRetrofit.serviceAPI.login_psswrd(new LoginRequest(phone, password)).enqueue(new Callback<LoginResult>() {
                        @Override
                        public void onResponse(Call<LoginResult> call, Response<LoginResult> response) {
                            binding.btnLoginPsswrd.setVisibility(View.VISIBLE);
                            binding.progressRegisterBarLogin.setVisibility(View.GONE);
                            //得到token，使用mmkv，本地持久化
                            if (response.isSuccessful()) {
                                MMKV mmkv = MMKV.defaultMMKV();
                                Log.d("ThereIsProblem", response.body().getData().getToken());
                                mmkv.encode("token", response.body().getData().getToken());
                                Toast.makeText(getActivity(), "登录成功", Toast.LENGTH_SHORT).show();
                                ARouter.getInstance().build("/sellcowhourse/app_MainActivity").withFlags(Intent.FLAG_ACTIVITY_NEW_TASK | Intent.FLAG_ACTIVITY_CLEAR_TASK).navigation();
                                getActivity().overridePendingTransition(com.example.common.R.anim.set_in, com.example.common.R.anim.set_out);
                            } else {
//                            Toast.makeText(getActivity(), response.body().getMsg(), Toast.LENGTH_SHORT).show();
                                Toast.makeText(getActivity(), "出问题啦", Toast.LENGTH_SHORT).show();
                            }
                        }

                        @Override
                        public void onFailure(Call<LoginResult> call, Throwable t) {
                            binding.btnLoginPsswrd.setVisibility(View.VISIBLE);
                            binding.progressRegisterBarLogin.setVisibility(View.GONE);
//                            ARouter.getInstance().build("/sellcowhourse/app_MainActivity").withFlags(Intent.FLAG_ACTIVITY_NEW_TASK | Intent.FLAG_ACTIVITY_CLEAR_TASK).navigation();
//                            getActivity().overridePendingTransition(R.anim.alpha_login_in, R.anim.alpha_loin_out);
//                            getActivity().overridePendingTransition(com.example.common.R.anim.set_in, com.example.common.R.anim.set_out);
                            Log.d("ThereIsProblem", t.toString());
                            Toast.makeText(getActivity(), "网络有问题呢", Toast.LENGTH_SHORT).show();
                        }
                    });
                } else {
                    Toast.makeText(getActivity(), "请完善信息", Toast.LENGTH_SHORT).show();
                }
            }
        });
        binding.chipPsswrd.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
//                SignInCode
//                Fragment signInCodeFragment = new SignInCodeFragment("登录", list);
//                TabLayout tabLayout = getActivity().findViewById(R.id.tl);
//                list.remove(0);
//                list.add(signInCodeFragment);
//                ViewPager viewPager = getActivity().findViewById(R.id.vp);
                FragmentManager supportFragmentManager = getActivity().getSupportFragmentManager();
                SignInCodeFragment signInCodeFragment = new SignInCodeFragment("登录");
                FragmentTransaction fragmentTransaction = supportFragmentManager.beginTransaction();
                fragmentTransaction.setCustomAnimations(android.R.anim.fade_in, android.R.anim.fade_out);
                fragmentTransaction.replace(R.id.frame_code, signInCodeFragment);
                fragmentTransaction.commit();
                binding.framePsswrd.setVisibility(View.GONE);
            }
        });
        return view;
    }
}