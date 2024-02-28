package com.example.mine.fragment.listviewfragment.productfragment;

import android.os.Bundle;
import android.os.Handler;
import android.os.HandlerThread;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Toast;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.fragment.app.FragmentManager;
import androidx.fragment.app.FragmentTransaction;
import androidx.recyclerview.widget.GridLayoutManager;
import androidx.swiperefreshlayout.widget.SwipeRefreshLayout;

import com.example.common.room.AppDatabase;
import com.example.common.room.entitues.User;
import com.example.core_net_work.MyRetrofit;
import com.example.core_net_work.model.goods.ProductSimpleInfoResult;
import com.example.core_net_work.model.goods.UserLoadedResult;
import com.example.mine.R;
import com.example.mine.databinding.FragmentProductSellingBinding;
import com.example.mine.fragment.listviewfragment.AddProductFragment;
import com.example.mine.fragment.listviewfragment.adapter.ProductRecyclerViewAdapter;
import com.example.mine.fragment.listviewfragment.entity.ProductEntity;

import java.util.ArrayList;
import java.util.List;
import java.util.function.Consumer;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

/**
 * @Author winiymissl
 * @Date 2024-02-25 22:26
 * @Version 1.0
 */
public class ProductSellIngFragment extends Fragment {
    FragmentProductSellingBinding binding;
    String name;

    public ProductSellIngFragment(String name) {
        this.name = name;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    List<ProductEntity> list = new ArrayList();
    ProductRecyclerViewAdapter adapter;

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        binding = FragmentProductSellingBinding.bind(inflater.inflate(R.layout.fragment_product_selling, container, false));
        binding.transparentViewCollect.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {

            }
        });
        requestData();
        binding.productSellingRecyclerview.setLayoutManager(new GridLayoutManager(getActivity(), 1));
        adapter = new ProductRecyclerViewAdapter(list, getActivity());
        binding.productSellingRecyclerview.setAdapter(adapter);
        binding.productSellingRefresh.setOnRefreshListener(new SwipeRefreshLayout.OnRefreshListener() {
            @Override
            public void onRefresh() {
                //清空重新加载
                refreshData();
            }
        });
        binding.floatingButtonAdd.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                FragmentManager fragmentManager = getActivity().getSupportFragmentManager();
                FragmentTransaction fragmentTransaction = fragmentManager.beginTransaction();
                AddProductFragment fragment = new AddProductFragment();
                fragmentTransaction.setCustomAnimations(com.example.common.R.anim.slide_in_from_right, com.example.common.R.anim.slide_out_to_right, com.example.common.R.anim.slide_in_from_right, com.example.common.R.anim.slide_out_to_right);
                fragmentTransaction.add(R.id.frame_mine, fragment);
                fragmentTransaction.addToBackStack(null);
                fragmentTransaction.commit();
            }
        });
        return binding.getRoot();
    }

    private void refreshData() {
        HandlerThread handlerThread = new HandlerThread("handlerThread");
        handlerThread.start();
        Handler handler = new Handler(handlerThread.getLooper());
        handler.post(new Runnable() {
            @Override
            public void run() {
                binding.productSellingConstrain.setVisibility(View.GONE);
                binding.textViewNone.setVisibility(View.GONE);
                List<User> allInfo = AppDatabase.getInstance(getActivity()).userDao().getAllInfo();
                User user = allInfo.get(0);
                MyRetrofit.serviceAPI.getUserUploaded(user.getUser_identity()).enqueue(new Callback<UserLoadedResult>() {
                    @Override
                    public void onResponse(Call<UserLoadedResult> call, Response<UserLoadedResult> response) {
                        if (response.isSuccessful()) {
                            list.clear();
                            response.body().getData().getUser_products().getCommodity().forEach(new Consumer<ProductSimpleInfoResult.CommodityInfo>() {
                                @Override
                                public void accept(ProductSimpleInfoResult.CommodityInfo commodityInfo) {
                                    list.add(new ProductEntity(commodityInfo.getMedia().get(0).getImage(), commodityInfo.getTitle(), String.valueOf(commodityInfo.getPrice())));
                                    adapter.notifyDataSetChanged();
                                }
                            });
                        }
                        binding.productSellingRefresh.setRefreshing(false);
                        binding.productSellingConstrain.setVisibility(View.VISIBLE);
                        binding.progressbarSellingLoading.setVisibility(View.GONE);
                    }

                    @Override
                    public void onFailure(Call<UserLoadedResult> call, Throwable t) {
                        Toast.makeText(getActivity(), "加载失败", Toast.LENGTH_SHORT).show();
                        binding.productSellingRefresh.setRefreshing(false);
                        binding.productSellingConstrain.setVisibility(View.VISIBLE);
                        binding.progressbarSellingLoading.setVisibility(View.GONE);
                    }
                });
            }
        });
    }

    private void requestData() {
        HandlerThread handlerThread = new HandlerThread("handlerThread");
        handlerThread.start();
        Handler handler = new Handler(handlerThread.getLooper());
        handler.post(new Runnable() {
            @Override
            public void run() {
                binding.textViewNone.setVisibility(View.GONE);
                binding.progressbarSellingLoading.setVisibility(View.VISIBLE);
                binding.productSellingConstrain.setVisibility(View.GONE);
                binding.productSellingRecyclerview.setVisibility(View.GONE);
                List<User> allInfo = AppDatabase.getInstance(getActivity()).userDao().getAllInfo();
                User user = allInfo.get(0);
                MyRetrofit.serviceAPI.getUserUploaded(user.getUser_identity()).enqueue(new Callback<UserLoadedResult>() {
                    @Override
                    public void onResponse(Call<UserLoadedResult> call, Response<UserLoadedResult> response) {
                        if (response.isSuccessful()) {
                            response.body().getData().getUser_products().getCommodity().forEach(new Consumer<ProductSimpleInfoResult.CommodityInfo>() {
                                @Override
                                public void accept(ProductSimpleInfoResult.CommodityInfo commodityInfo) {
//                                    list.clear();
                                    list.add(new ProductEntity(commodityInfo.getMedia().get(0).getImage(), commodityInfo.getTitle(), String.valueOf(commodityInfo.getPrice())));
                                    adapter.notifyDataSetChanged();
                                }
                            });
                            if (list.size() == 0) {
                                binding.textViewNone.setVisibility(View.VISIBLE);
                            }
                        }
                        binding.productSellingConstrain.setVisibility(View.VISIBLE);
                        binding.productSellingRecyclerview.setVisibility(View.VISIBLE);
                        binding.progressbarSellingLoading.setVisibility(View.GONE);
                    }

                    @Override
                    public void onFailure(Call<UserLoadedResult> call, Throwable t) {
                        binding.textViewNone.setVisibility(View.VISIBLE);
                        Toast.makeText(getActivity(), "加载失败", Toast.LENGTH_SHORT).show();
                        binding.productSellingRecyclerview.setVisibility(View.VISIBLE);
                        binding.productSellingConstrain.setVisibility(View.VISIBLE);
                        binding.progressbarSellingLoading.setVisibility(View.GONE);
                    }
                });
            }
        });
    }
}
