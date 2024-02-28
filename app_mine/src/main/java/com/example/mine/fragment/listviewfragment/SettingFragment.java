package com.example.mine.fragment.listviewfragment;

import android.os.Bundle;
import android.util.Log;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;

import com.example.mine.R;
import com.example.mine.databinding.FragmentSettingBinding;

/**
 * @Author winiymissl
 * @Date 2024-02-23 16:55
 * @Version 1.0
 */
public class SettingFragment extends Fragment {
    FragmentSettingBinding binding;

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
//        try {
            binding = FragmentSettingBinding.bind(inflater.inflate(R.layout.fragment_setting, container, false));
            binding.transparentViewSetting.setOnClickListener(new View.OnClickListener() {
                @Override
                public void onClick(View v) {

                }
            });
//        } catch (Exception e) {
//            Log.d("这里有一个问题", e.toString());
//        }
        return binding.getRoot();
    }
}
