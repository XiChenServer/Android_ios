package com.example.mine.fragment.listviewfragment.productfragment;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Toast;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.recyclerview.widget.GridLayoutManager;
import androidx.swiperefreshlayout.widget.SwipeRefreshLayout;

import com.example.core_net_work.MyRetrofit;
import com.example.core_net_work.model.goods.BoughtResult;
import com.example.mine.R;
import com.example.mine.databinding.FragmentBoughtBinding;
import com.example.mine.fragment.listviewfragment.adapter.ProductRecyclerViewAdapter;
import com.example.mine.fragment.listviewfragment.entity.ProductEntity;
import com.tencent.mmkv.MMKV;

import java.util.ArrayList;
import java.util.List;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

/**
 * @Author winiymissl
 * @Date 2024-02-28 16:28
 * @Version 1.0
 */
public class BoughtFragment extends Fragment {
    String name;

    public BoughtFragment(String name) {
        this.name = name;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    FragmentBoughtBinding binding;

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_bought, container, false);
        binding = FragmentBoughtBinding.bind(view);
        binding.transparentViewBought.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {

            }
        });
        List<ProductEntity> list = new ArrayList<>();
        ProductRecyclerViewAdapter adapter = new ProductRecyclerViewAdapter(list, getActivity());
        binding.recyclerviewBought.setLayoutManager(new GridLayoutManager(getActivity(), 1));
        binding.recyclerviewBought.setAdapter(adapter);
        requestData();
        binding.refreshBought.setOnRefreshListener(new SwipeRefreshLayout.OnRefreshListener() {
            @Override
            public void onRefresh() {
                refreshData();
            }
        });
        return binding.getRoot();
    }

    private void refreshData() {
        binding.progressBought.setVisibility(View.VISIBLE);
        binding.textViewBought.setVisibility(View.GONE);

        binding.constraintBought.setVisibility(View.GONE);
        MyRetrofit.serviceAPI.getBoughtInfo("Bearer " + MMKV.defaultMMKV().getString("token", null)).enqueue(new Callback<BoughtResult>() {
            @Override
            public void onResponse(Call<BoughtResult> call, Response<BoughtResult> response) {
                if (response.isSuccessful()) {
                    if (response.body().getOrders().size() == 0) {
                        binding.textViewBought.setVisibility(View.VISIBLE);
                    }
                } else {
                    binding.textViewBought.setText("加载有问题");
                    binding.textViewBought.setVisibility(View.VISIBLE);
                }
                binding.constraintBought.setVisibility(View.VISIBLE);
                binding.progressBought.setVisibility(View.GONE);
                binding.refreshBought.setRefreshing(false);
            }

            @Override
            public void onFailure(Call<BoughtResult> call, Throwable t) {
                binding.textViewBought.setText("加载失败");
                binding.textViewBought.setVisibility(View.VISIBLE);
                Toast.makeText(getActivity(), "加载失败", Toast.LENGTH_SHORT).show();
                binding.progressBought.setVisibility(View.GONE);
                binding.constraintBought.setVisibility(View.VISIBLE);
                binding.refreshBought.setRefreshing(false);
            }
        });
    }

    private void requestData() {
        binding.progressBought.setVisibility(View.VISIBLE);
        binding.constraintBought.setVisibility(View.GONE);
        binding.textViewBought.setVisibility(View.GONE);
        MyRetrofit.serviceAPI.getBoughtInfo("Bearer " + MMKV.defaultMMKV().getString("token", null)).enqueue(new Callback<BoughtResult>() {
            @Override
            public void onResponse(Call<BoughtResult> call, Response<BoughtResult> response) {
                if (response.isSuccessful()) {
                    if (response.body().getOrders().size() == 0) {
                        binding.textViewBought.setVisibility(View.VISIBLE);
                    }
                } else {
                    binding.textViewBought.setText("加载有问题");
                    binding.textViewBought.setVisibility(View.VISIBLE);
                }
                binding.constraintBought.setVisibility(View.VISIBLE);
                binding.progressBought.setVisibility(View.GONE);
            }

            @Override
            public void onFailure(Call<BoughtResult> call, Throwable t) {
                binding.textViewBought.setText("加载失败");
                binding.textViewBought.setVisibility(View.VISIBLE);
                Toast.makeText(getActivity(), "加载失败", Toast.LENGTH_SHORT).show();
                binding.progressBought.setVisibility(View.GONE);
                binding.constraintBought.setVisibility(View.VISIBLE);
            }
        });
    }
}
