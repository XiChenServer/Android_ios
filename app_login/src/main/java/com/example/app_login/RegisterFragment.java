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
                        binding.register.setVisibility(View.GONE);
                        binding.progressRegisterBar.setVisibility(View.VISIBLE);
                        MyRetrofit.serviceAPI.register(new RegisterRequest(binding.etPhone.getText().toString(), binding.etPsswrd.getText().toString(), binding.etCode.getText().toString())).enqueue(new Callback<RegisterResult>() {
                            @Override
                            public void onResponse(Call<RegisterResult> call, Response<RegisterResult> response) {
                                binding.progressRegisterBar.setVisibility(View.GONE);
                                binding.register.setVisibility(View.VISIBLE);
                                if (response.isSuccessful()) {
                                    Log.d("IsThereProblem", response.body().getMsg());
                                    Toast.makeText(getActivity(), "注册成功:", Toast.LENGTH_SHORT).show();
                                } else {
                                    Toast.makeText(getActivity(), "出问题了呢", Toast.LENGTH_SHORT).show();
                                    binding.etPhone.setText("");
                                    binding.etCode.setText("");
                                    binding.etPsswrdToo.setText("");
                                    binding.etPsswrd.setText("");
                                }
                            }

                            @Override
                            public void onFailure(Call<RegisterResult> call, Throwable t) {
                                binding.progressRegisterBar.setVisibility(View.GONE);
                                binding.register.setVisibility(View.VISIBLE);
                                Log.d("IsThereProblem", t.toString());
                                Toast.makeText(getActivity(), "错误", Toast.LENGTH_SHORT).show();
                                binding.etPhone.setText("");
                                binding.etCode.setText("");
                                binding.etPsswrdToo.setText("");
                                binding.etPsswrd.setText("");
                            }
                        });
                    } else {
                        Toast.makeText(getActivity(), "请完善信息", Toast.LENGTH_SHORT).show();
                    }
                }
            });
            binding.btnCode.setOnClickListener(new View.OnClickListener() {
                @Override
                public void onClick(View v) {

                    if (!binding.etPhone.getText().toString().equals("")) {
                        binding.btnCode.setVisibility(View.GONE);
                        binding.progressBar.setVisibility(View.VISIBLE);
                        Log.d("IsThereProblem", binding.etPhone.getText().toString());
                        binding.btnCode.setEnabled(false);

                        MyRetrofit.serviceAPI.getCode(new CodeRequest(binding.etPhone.getText().toString())).enqueue(new Callback<CodeResult>() {
                            @Override
                            public void onResponse(Call<CodeResult> call, Response<CodeResult> response) {
                                binding.progressBar.setVisibility(View.GONE);
                                binding.btnCode.setVisibility(View.VISIBLE);
//                                Toast.makeText(getActivity(), "Code : " + String.valueOf(response.body().getCode()), Toast.LENGTH_SHORT).show();
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
                                        binding.btnCode.setText("验证码");
                                    }
                                }.start();
                                Toast.makeText(getActivity(), "msg : " + String.valueOf(response.body().getMsg()), Toast.LENGTH_SHORT).show();
                            }

                            @Override
                            public void onFailure(Call<CodeResult> call, Throwable t) {
                                Toast.makeText(getActivity(), "网络崩溃了呢", Toast.LENGTH_SHORT).show();
//                                Log.d("IsThereProblem", t.toString());
                                binding.progressBar.setVisibility(View.GONE);
                                binding.btnCode.setVisibility(View.VISIBLE);
                                binding.btnCode.setEnabled(true);
//                                binding.btnCode.setText("验证码");
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
