package com.example.mine.fragment.listviewfragment.productfragment;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.recyclerview.widget.GridLayoutManager;
import androidx.swiperefreshlayout.widget.SwipeRefreshLayout;

import com.example.core_net_work.MyRetrofit;
import com.example.core_net_work.model.goods.BoughtResult;
import com.example.mine.R;
import com.example.mine.databinding.FragmentSellOutBinding;
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
 * @Date 2024-02-28 16:29
 * @Version 1.0
 */
public class SellOutOrderFragment extends Fragment {
    private String name;

    public SellOutOrderFragment(String name) {
        this.name = name;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    FragmentSellOutBinding binding;

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_sell_out, container, false);
        binding = FragmentSellOutBinding.bind(view);
        binding.transparentViewSellOut.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {

            }
        });
        List<ProductEntity> list = new ArrayList<>();
        ProductRecyclerViewAdapter adapter = new ProductRecyclerViewAdapter(list, getActivity());
        binding.recyclerviewSellOut.setLayoutManager(new GridLayoutManager(getActivity(), 1));
        binding.recyclerviewSellOut.setAdapter(adapter);
        requestData();
        binding.refreshSellOut.setOnRefreshListener(new SwipeRefreshLayout.OnRefreshListener() {
            @Override
            public void onRefresh() {
                refreshData();
            }
        });
        return view;
    }

    private void refreshData() {
        binding.constraintSellOut.setVisibility(View.GONE);
        binding.textViewSellOut.setVisibility(View.GONE);
        MyRetrofit.serviceAPI.getBoughtInfo("Bearer "+MMKV.defaultMMKV().getString("token", null)).enqueue(new Callback<BoughtResult>() {
            @Override
            public void onResponse(Call<BoughtResult> call, Response<BoughtResult> response) {
                if (response.isSuccessful()) {
                    if (response.body().getOrders().size() == 0) {
                        binding.textViewSellOut.setVisibility(View.VISIBLE);
                    }
                } else {
                    binding.textViewSellOut.setText("加载有问题");
                    binding.textViewSellOut.setVisibility(View.VISIBLE);
                }
                binding.constraintSellOut.setVisibility(View.VISIBLE);
                binding.progressSellOut.setVisibility(View.GONE);
                binding.refreshSellOut.setRefreshing(false);
            }

            @Override
            public void onFailure(Call<BoughtResult> call, Throwable t) {
                binding.textViewSellOut.setText("加载失败");
                binding.textViewSellOut.setVisibility(View.VISIBLE);
                binding.progressSellOut.setVisibility(View.GONE);
                binding.constraintSellOut.setVisibility(View.VISIBLE);
                binding.refreshSellOut.setRefreshing(false);
            }
        });
    }

    private void requestData() {
        binding.constraintSellOut.setVisibility(View.GONE);
        binding.textViewSellOut.setVisibility(View.GONE);
        MyRetrofit.serviceAPI.getBoughtInfo("Bearer "+MMKV.defaultMMKV().getString("token", null)).enqueue(new Callback<BoughtResult>() {
            @Override
            public void onResponse(Call<BoughtResult> call, Response<BoughtResult> response) {
                if (response.isSuccessful()) {
                    if (response.body().getOrders().size() == 0) {
                        binding.textViewSellOut.setVisibility(View.VISIBLE);
                    }
                } else {
                    binding.textViewSellOut.setText("加载有问题");
                    binding.textViewSellOut.setVisibility(View.VISIBLE);
                }
                binding.constraintSellOut.setVisibility(View.VISIBLE);
                binding.progressSellOut.setVisibility(View.GONE);
                binding.refreshSellOut.setRefreshing(false);
            }

            @Override
            public void onFailure(Call<BoughtResult> call, Throwable t) {
                binding.textViewSellOut.setText("加载失败");
                binding.textViewSellOut.setVisibility(View.VISIBLE);
                binding.progressSellOut.setVisibility(View.GONE);
                binding.constraintSellOut.setVisibility(View.VISIBLE);
                binding.refreshSellOut.setRefreshing(false);
            }
        });
    }
}
