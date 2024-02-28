package com.example.mine.fragment.listviewfragment;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;

import com.example.mine.R;
import com.example.mine.databinding.FragmentCollectBinding;

/**
 * @Author winiymissl
 * @Date 2024-02-23 16:53
 * @Version 1.0
 */
public class CollectionFragment extends Fragment {
    FragmentCollectBinding binding;

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        binding = FragmentCollectBinding.bind(inflater.inflate(R.layout.fragment_collect, container, false));


        binding.transparentViewCollect.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {

            }
        });
        return binding.getRoot();
    }
}
