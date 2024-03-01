package com.example.mine_pager.fragment;

import android.os.Bundle;
import android.os.Handler;
import android.os.HandlerThread;
import android.util.Log;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Toast;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.appcompat.app.AppCompatActivity;
import androidx.fragment.app.Fragment;
import androidx.fragment.app.FragmentManager;
import androidx.fragment.app.FragmentTransaction;

import com.example.common.room.AppDatabase;
import com.example.common.room.entitues.ShoppingCarOrder;
import com.example.core_net_work.MyRetrofit;
import com.example.core_net_work.model.goods.ProductInfoResult;
import com.example.core_net_work.model.goods.ProductSimpleInfoResult;
import com.example.mine_pager.R;
import com.example.mine_pager.adapter.MyBannerAdapter;
import com.example.mine_pager.databinding.FragmentDetailBinding;
import com.google.android.material.appbar.MaterialToolbar;
import com.youth.banner.indicator.CircleIndicator;
import com.youth.banner.transformer.ZoomOutPageTransformer;

import java.util.ArrayList;
import java.util.List;
import java.util.function.Consumer;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

/**
 * @Author winiymissl
 * @Date 2024-02-29 9:52
 * @Version 1.0
 */
public class DetailFragment extends Fragment {
    FragmentDetailBinding binding;
    List<DataBean> list = new ArrayList<>();

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_detail, container, false);
        binding = FragmentDetailBinding.bind(view);
        Bundle bundle = getArguments();
        String id = bundle.getString("id", null);
        MaterialToolbar toolbar = binding.materialToolbar;
        ((AppCompatActivity) requireActivity()).setSupportActionBar(toolbar);
        ((AppCompatActivity) requireActivity()).getSupportActionBar().setDisplayHomeAsUpEnabled(true);
        binding.materialToolbar.setTitle("");

        toolbar.setNavigationOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                FragmentManager fragmentManager = getActivity().getSupportFragmentManager();
                FragmentTransaction fragmentTransaction = fragmentManager.beginTransaction();
                fragmentManager.popBackStack();
                fragmentTransaction.commit();
            }
        });

        binding.chipAdd.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                binding.counter.setText(String.valueOf(Integer.valueOf(binding.counter.getText().toString()) + 1));
            }
        });
        binding.chipMinus.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                int count = Integer.valueOf(binding.counter.getText().toString());
                if (count > 0) {
                    binding.counter.setText(String.valueOf(count - 1));
                }
            }
        });


        binding.banner.setPageTransformer(new ZoomOutPageTransformer());
        MyRetrofit.serviceAPI.getProductInfo(id).enqueue(new Callback<ProductInfoResult>() {
            @Override
            public void onResponse(Call<ProductInfoResult> call, Response<ProductInfoResult> response) {
                if (response.isSuccessful()) {
                    ProductInfoResult.ProductCommodity.ProductOneInfo oneInfo = response.body().getData().getCommodity();
                    response.body().getData().getCommodity().getMedia().forEach(new Consumer<ProductSimpleInfoResult.CommodityInfo.MediaResult>() {
                        @Override
                        public void accept(ProductSimpleInfoResult.CommodityInfo.MediaResult mediaResult) {
                            list.add(new DataBean(mediaResult.getImage()));
                        }
                    });
                    Log.d("这是一个问题", list.toString());
                    MyBannerAdapter adapter = new MyBannerAdapter(list, getActivity());
                    binding.banner.addBannerLifecycleObserver(getActivity())//添加生命周期观察者
                            .setAdapter(adapter).setIndicator(new CircleIndicator(getActivity()));
//                    binding.textViewInformation.setText(oneInfo.getInformation());
                    binding.detailInclude.textViewPriceDetail.setText("￥" + oneInfo.getPrice());
                    binding.detailInclude.textViewNameDetail.setText(oneInfo.getTitle());
                    binding.detailInclude.textViewInfoDetail.setText(oneInfo.getInformation());
                    binding.floatingActionButtonCreate.setOnClickListener(new View.OnClickListener() {
                        @Override
                        public void onClick(View v) {
                            HandlerThread handlerThread = new HandlerThread("save_shopping_car");
                            handlerThread.start();
                            Handler handler = new Handler(handlerThread.getLooper());
                            handler.post(new Runnable() {
                                @Override
                                public void run() {
                                    List<ShoppingCarOrder> temp = new ArrayList<>();
                                    temp.add(new ShoppingCarOrder(oneInfo.getMedia().get(0).getImage(), Float.valueOf(oneInfo.getPrice()), Integer.valueOf(binding.counter.getText().toString()), oneInfo.getTitle()));
                                    AppDatabase.getInstance(getActivity()).shoppingCarDao().insertAll(temp);
                                    Toast.makeText(getActivity(), "添加成功", Toast.LENGTH_SHORT).show();
//                                    Snackbar.make(view, "添加成功!", Snackbar.LENGTH_SHORT).setAction("去看看", new View.OnClickListener() {
//                                        @Override
//                                        public void onClick(View v) {
//                                        }
//                                    }).show();
                                }
                            });

                        }
                    });
                }
            }

            @Override
            public void onFailure(Call<ProductInfoResult> call, Throwable t) {
                Toast.makeText(getActivity(), "返回失败", Toast.LENGTH_SHORT).show();
            }
        });
        return view;
    }
}
