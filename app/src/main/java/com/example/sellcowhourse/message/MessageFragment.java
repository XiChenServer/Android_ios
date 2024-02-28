package com.example.sellcowhourse.message;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;

import com.example.sellcowhourse.R;
import com.example.sellcowhourse.databinding.FragmentMessageBinding;

/**
 * @Author winiymissl
 * @Date 2024-02-20 21:00
 * @Version 1.0
 */
public class MessageFragment extends Fragment {

    FragmentMessageBinding binding;

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_message, container, false);
        binding = FragmentMessageBinding.bind(view);

        return view;
    }
}
