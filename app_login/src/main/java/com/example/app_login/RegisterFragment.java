package com.example.app_login;

import android.os.Bundle;
import android.os.CountDownTimer;
import android.text.Editable;
import android.text.TextWatcher;
import android.util.Log;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Toast;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;

import com.example.app_login.databinding.FragmentRegisterBinding;
import com.example.core_net_work.MyRetrofit;
import com.example.core_net_work.model.login.CodeRequest;
import com.example.core_net_work.model.login.CodeResult;
import com.example.core_net_work.model.login.RegisterRequest;
import com.example.core_net_work.model.login.RegisterResult;
import com.google.android.material.tabs.TabLayout;
import com.google.android.material.textfield.TextInputLayout;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

/**
 * @Author winiymissl
 * @Date 2023-12-11 14:12
 * @Version 1.0
 */
public class RegisterFragment extends Fragment {
    FragmentRegisterBinding binding;
    String name;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public RegisterFragment(String name) {
        this.name = name;
    }

    public RegisterFragment() {

    }

    private CountDownTimer countDownTimer;

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {

        View view = null;
        try {

            view = inflater.inflate(R.layout.fragment_register, container, false);
            binding = FragmentRegisterBinding.bind(view);
//            binding.register.setEnabled(false);
//            binding.etCode.addTextChangedListener(new TextWatcher() {
//                @Override
//                public void beforeTextChanged(CharSequence s, int start, int count, int after) {
//
//                }
//
//                @Override
//                public void onTextChanged(CharSequence s, int start, int before, int count) {
//                    binding.register.setEnabled(true);
//                }
//
//                @Override
//                public void afterTextChanged(Editable s) {
//
//                }
//            });
            binding.register.setOnClickListener(new View.OnClickListener() {
                @Override
                public void onClick(View v) {
                    //对信息进行判断
                    if (!binding.etCode.getText().toString().equals("") && !binding.etPsswrd.getText().toString().equals("") && !binding.etPsswrdToo.getText().toString().equals("") && !binding.etPhone.getText().toString().equals("") && binding.etPsswrdToo.getText().toString().equals(binding.etPsswrdToo.getText().toString())) {
                        MyRetrofit.serviceAPI.register(new RegisterRequest(binding.etPhone.getText().toString(), binding.etPsswrd.getText().toString(), String.valueOf(binding.btnCode).toString())).enqueue(new Callback<RegisterResult>() {
                            @Override
                            public void onResponse(Call<RegisterResult> call, Response<RegisterResult> response) {
//                                if (response.isSuccessful()) {
                                if (response.body() == null) {
                                    Log.d("IsThereProblem", "返回值为空");
                                } else {
                                    Log.d("IsThereProblem", "返回值不为空");
                                }
                                Toast.makeText(getActivity(), "注册成功:", Toast.LENGTH_SHORT).show();
                                Log.d("IsThereProblem", response.body().getMsg());
                                TabLayout tabLayout = getActivity().findViewById(R.id.tl);
                                TabLayout.Tab tab = tabLayout.getTabAt(0);
                                tab.select();
//                                }
                            }

                            @Override
                            public void onFailure(Call<RegisterResult> call, Throwable t) {
                                Log.d("IsThereProblem", t.toString());
                                Toast.makeText(getActivity(), "注册失败", Toast.LENGTH_SHORT).show();
                            }
                        });
                    } else {
                        Toast.makeText(getActivity(), "信息有误", Toast.LENGTH_SHORT).show();
                    }
                }
            });
            binding.btnCode.setOnClickListener(new View.OnClickListener() {
                @Override
                public void onClick(View v) {
                    if (!binding.etPhone.getText().toString().equals("")) {
                        Log.d("IsThereProblem", binding.etPhone.getText().toString());
                        binding.btnCode.setEnabled(false);
                        countDownTimer = new CountDownTimer(60000, 1000) {
                            @Override
                            public void onTick(long millisUntilFinished) {
                                // 更新按钮上显示的文本，显示剩余时间
                                binding.btnCode.setText("(" + millisUntilFinished / 1000 + "秒)");
                            }

                            @Override
                            public void onFinish() {
                                // 倒计时结束，启用按钮并重置文本
                                binding.btnCode.setEnabled(true);
                                binding.btnCode.setText("获取验证码");
                            }
                        }.start();
                        MyRetrofit.serviceAPI.getCode(new CodeRequest(binding.etPhone.getText().toString())).enqueue(new Callback<CodeResult>() {
                            @Override
                            public void onResponse(Call<CodeResult> call, Response<CodeResult> response) {
//                                Toast.makeText(getActivity(), "Code : " + String.valueOf(response.body().getCode()), Toast.LENGTH_SHORT).show();
                                Toast.makeText(getActivity(), "msg : " + String.valueOf(response.body().getMsg()), Toast.LENGTH_SHORT).show();
                            }

                            @Override
                            public void onFailure(Call<CodeResult> call, Throwable t) {
                                Log.d("IsThereProblem", t.toString());
                            }
                        });
                    } else {
                        Toast.makeText(getActivity(), "请填写手机号码", Toast.LENGTH_SHORT).show();
                    }
                }
            });
        } catch (Exception e) {
            Log.d("IsThereProblem", e.toString());
        }
        return view;
    }

    @Override
    public void onDestroyView() {
        super.onDestroyView();
        // 释放 ViewBinding 对象
        binding = null;
    }
}
