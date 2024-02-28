package com.example.mine.fragment;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Toast;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;

import com.example.core_net_work.MyRetrofit;
import com.example.core_net_work.model.login.CodeRequest;
import com.example.core_net_work.model.login.CodeResult;
import com.example.core_net_work.model.userInfo.ChangeUserInfoRequest;
import com.example.core_net_work.model.userInfo.ChangeUserInfoResult;
import com.example.mine.R;
import com.example.mine.databinding.FragmentAppMineSaveinfoBinding;
import com.tencent.mmkv.MMKV;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

/**
 * @Author winiymissl
 * @Date 2024-02-20 19:27
 * @Version 1.0
 */
public class app_mine_saveInfoFragment extends Fragment {
    FragmentAppMineSaveinfoBinding binding;

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_app_mine_saveinfo, container, false);
        binding = FragmentAppMineSaveinfoBinding.bind(view);
        Bundle bundle = getArguments();
        String nickname = bundle.getString("nickname");
        String phone = bundle.getString("phone");
        binding.buttonGetCode.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                MyRetrofit.serviceAPI.getCode(new CodeRequest(phone)).enqueue(new Callback<CodeResult>() {
                    @Override
                    public void onResponse(Call<CodeResult> call, Response<CodeResult> response) {
                        if (response.isSuccessful()) {
                            Toast.makeText(getActivity(), "发送成功", Toast.LENGTH_SHORT).show();
                        }

                    }

                    @Override
                    public void onFailure(Call<CodeResult> call, Throwable t) {
                        Toast.makeText(getActivity(), "发送失败", Toast.LENGTH_SHORT).show();
                    }
                });
            }
        });
//        binding.buttonSubmit.setOnClickListener(new View.OnClickListener() {
//            @Override
//            public void onClick(View v) {
//                MyRetrofit.serviceAPI.changeUserInfo(MMKV.defaultMMKV().getString("token", null),
//                                new ChangeUserInfoRequest(nickname),
//                                binding.inputTextCode.getText().toString())
//                        .enqueue(new Callback<ChangeUserInfoResult>() {
//                            @Override
//                            public void onResponse(Call<ChangeUserInfoResult> call, Response<ChangeUserInfoResult> response) {
//                                if (response.isSuccessful()) {
//                                    Toast.makeText(getActivity(), "更新成功", Toast.LENGTH_SHORT).show();
//
//                                } else {
//                                    Toast.makeText(getActivity(), "更新失败", Toast.LENGTH_SHORT).show();
//                                }
//                            }
//
//                            @Override
//                            public void onFailure(Call<ChangeUserInfoResult> call, Throwable t) {
//                                Toast.makeText(getActivity(), "更新失败", Toast.LENGTH_SHORT).show();
//                            }
//                        });
//            }
//        });
        return view;
    }
}
