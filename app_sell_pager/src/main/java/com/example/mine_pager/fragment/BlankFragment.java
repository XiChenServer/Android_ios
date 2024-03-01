package com.example.mine_pager.fragment;

import android.os.Bundle;
import android.os.Handler;
import android.os.HandlerThread;
import android.util.Log;
import android.view.LayoutInflater;
import android.view.MenuItem;
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

import com.alibaba.android.arouter.launcher.ARouter;
import com.example.common.room.AppDatabase;
import com.example.common.room.dao.ProductSimpleDao;
import com.example.common.room.entitues.ProductSimple;
import com.example.common.room.recyclerviewitem.MyRecyclerViewItemClick;
import com.example.core_net_work.MyRetrofit;
import com.example.core_net_work.model.goods.ProductSimpleInfoResult;
import com.example.mine_pager.R;
import com.example.mine_pager.adapter.RecyclerViewAdapter;
import com.example.mine_pager.databinding.FragmentBlankBinding;
import com.kennyc.bottomsheet.BottomSheetListener;
import com.kennyc.bottomsheet.BottomSheetMenuDialogFragment;
import com.tencent.mmkv.MMKV;

import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;
import java.util.function.Consumer;

import jp.wasabeef.recyclerview.animators.FadeInRightAnimator;
import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;


public class BlankFragment extends Fragment {
    FragmentBlankBinding binding;

    private List<ProductSimple> list = new ArrayList();

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_blank, container, false);
        binding = FragmentBlankBinding.bind(view);
//        list.add(new ProductSimple(String.valueOf(com.example.common.R.drawable.loading), "Loading", "loading", "loading", "0"));
//        list.add(new ProductSimple(String.valueOf(com.example.common.R.drawable.loading), "Loading", "loading", "loading", "0"));
//        list.add(new ProductSimple(String.valueOf(com.example.common.R.drawable.loading), "Loading", "loading", "loading", "0"));
//        list.add(new ProductSimple(String.valueOf(com.example.common.R.drawable.loading), "Loading", "loading", "loading", "0"));
//         创建 Handler 对象，并将 Looper 对象传入
        RecyclerViewAdapter adapter = new RecyclerViewAdapter(list, getActivity());
        GridLayoutManager gridLayoutManager = new GridLayoutManager(getActivity(), 2);
        binding.sellPagerRecyclerView.setLayoutManager(gridLayoutManager);
        binding.sellPagerRecyclerView.setAdapter(adapter);
        //设置recyclerview的动画
        binding.sellPagerRecyclerView.setItemAnimator(new FadeInRightAnimator());
        binding.sellPagerRecyclerView.addOnItemTouchListener(new MyRecyclerViewItemClick(getActivity(), new MyRecyclerViewItemClick.OnItemClickListener() {
            @Override
            public void onItemClick(View view, int position) {
                FragmentManager fragmentManager = getActivity().getSupportFragmentManager();
                FragmentTransaction fragmentTransaction = fragmentManager.beginTransaction();
                DetailFragment fragment = new DetailFragment();
                Bundle bundle = new Bundle();
                bundle.putString("id", list.get(position).getId());
                Log.d("这是一个问题", list.get(position).getId());
                fragment.setArguments(bundle);
                fragmentTransaction.setCustomAnimations(com.example.common.R.anim.slide_in_from_right, com.example.common.R.anim.slide_out_to_right, com.example.common.R.anim.slide_in_from_right, com.example.common.R.anim.slide_out_to_right);
                fragmentTransaction.add(R.id.sell_pager_coordinatorlayout, fragment);
                fragmentTransaction.addToBackStack(null);
                fragmentTransaction.commit();
            }
        }));


        HandlerThread handlerThread = new HandlerThread("MyHandlerThread");
        handlerThread.start();
        Handler handler = new Handler(handlerThread.getLooper());
        handler.post(new Runnable() {
            @Override
            public void run() {
                List<ProductSimple> allInfo = AppDatabase.getInstance(getActivity()).productSimpleDao().getAllInfo();
                list.clear();
                list.addAll(allInfo);
                adapter.notifyDataSetChanged();
            }
        });
        Log.d("这是一个问题", MMKV.defaultMMKV().getString("token", null));
        binding.swipeRefreshLayoutSellPager.setOnRefreshListener(new SwipeRefreshLayout.OnRefreshListener() {
            @Override
            public void onRefresh() {
                MyRetrofit.serviceAPI.getProductSimpleInfo().enqueue(new Callback<ProductSimpleInfoResult>() {
                    @Override
                    public void onResponse(Call<ProductSimpleInfoResult> call, Response<ProductSimpleInfoResult> response) {
                        if (response.isSuccessful()) {
                            Toast.makeText(getActivity(), "成功", Toast.LENGTH_SHORT).show();
                            list.clear();
                            response.body().getData().getCommodity().forEach(new Consumer<ProductSimpleInfoResult.CommodityInfo>() {
                                @Override
                                public void accept(ProductSimpleInfoResult.CommodityInfo commodityInfo) {
                                    List<ProductSimpleInfoResult.CommodityInfo.MediaResult> media = commodityInfo.getMedia();
                                    list.add(new ProductSimple(commodityInfo.getCommodity_identity(), media.get(0).getImage(), commodityInfo.getTitle(), commodityInfo.getInformation(), String.valueOf(commodityInfo.getPrice()), "3.5", String.valueOf(commodityInfo.getID())));
                                    binding.swipeRefreshLayoutSellPager.setRefreshing(false);
                                    adapter.notifyDataSetChanged();
                                }
                            });
                            HandlerThread handlerThread = new HandlerThread("saveProductSimple");
                            handlerThread.start();
                            Handler handler = new Handler(handlerThread.getLooper());
                            handler.post(new Runnable() {
                                @Override
                                public void run() {
                                    ProductSimpleDao simpleDao = AppDatabase.getInstance(getActivity()).productSimpleDao();
                                    simpleDao.deleteAll();
                                    simpleDao.insertAll(list);
                                }
                            });
                        } else {
                            Toast.makeText(getActivity(), "有问题", Toast.LENGTH_SHORT).show();
                        }
                        binding.swipeRefreshLayoutSellPager.setRefreshing(false);
                    }

                    @Override
                    public void onFailure(Call<ProductSimpleInfoResult> call, Throwable t) {
                        binding.swipeRefreshLayoutSellPager.setRefreshing(false);
                        Toast.makeText(getActivity(), "失败了", Toast.LENGTH_SHORT).show();
                        Log.d("大问题", t.toString());
                    }
                });
            }
        });
//        binding.sellPagerRecyclerView.setOnRefreshListener(new OnRefreshListener() {
//            @Override
//            public void onRefresh() {

//        });

        //强制刷新
//        binding.sellPagerRecyclerView.setRefreshProgressStyle(ProgressStyle.Pacman);
//        binding.sellPagerRecyclerView.forceToRefresh();
//        binding.sellPagerRecyclerView.refreshComplete(pageSize);
//        adapter.notifyDataSetChanged();

//        binding.sellPagerRecyclerView.setLoadMoreEnabled(true);
//        binding.sellPagerRecyclerView.setLoadMoreFooter(new LoadingFooter(getActivity()), true);
//        binding.sellPagerRecyclerView.setFooterViewHint("拼命加载中", "已经全部为你呈现了", "网络不给力啊，点击再试一次吧");
//        binding.sellPagerRecyclerView.setOnLoadMoreListener(new OnLoadMoreListener() {
//            @Override
//            public void onLoadMore() {
//                Toast.makeText(getActivity(), "wertwert", Toast.LENGTH_SHORT).show();
//                Log.d("loadmore问题", String.valueOf(count));
//                binding.sellPagerRecyclerView.forceToRefresh();
//                count++;
//
//            }
//        });

//        adapter.setOnItemLongClickListener(new OnItemLongClickListener() {
//            @Override
//            public void onItemLongClick(View view, int position) {
//                new BottomSheetMenuDialogFragment.Builder(getActivity()).setSheet(R.menu.menu_more).setTitle("more").setListener(new BottomSheetListener() {
//                    @Override
//                    public void onSheetShown(@NonNull BottomSheetMenuDialogFragment bottomSheetMenuDialogFragment, @Nullable Object o) {
//
//                    }
//
//                    @Override
//                    public void onSheetItemSelected(@NonNull BottomSheetMenuDialogFragment bottomSheetMenuDialogFragment, @NonNull MenuItem menuItem, @Nullable Object o) {
//                        if (menuItem.getItemId() == R.id.more_shopping_car) {
//                            HandlerThread shoppingCar = new HandlerThread("readToShoppingCar");
//                            shoppingCar.start();
//                            Handler thread = new Handler(shoppingCar.getLooper());
//                            thread.post(new Runnable() {
//                                @Override
//                                public void run() {
//                                    List<ShoppingCarOrder> temp = new ArrayList<>();
//                                    if (temp.size() > 0) {
//                                        String productIdentity = list.get(position).getProduct_identity();
////                                    //通过唯一标识，找到这个商品，具体在哪里
//                                        try {
//                                            List<Product> product = AppDatabase.getInstance(getActivity()).productDao().getProduct(productIdentity);
//                                            Product product1 = product.get(0);
//                                            temp.add(new ShoppingCarOrder(product1.getImage(), Integer.valueOf(product1.getPrice()), 1, product1.getTitle()));
//                                            AppDatabase.getInstance(getActivity()).shoppingCarDao().insertAll(temp);
//                                        } catch (NumberFormatException e) {
//
//                                            Log.d("问题", e.toString());
//                                        }
//
//                                        getActivity().runOnUiThread(new Runnable() {
//                                            @Override
//                                            public void run() {
//                                                Snackbar.make(view, "添加成功", Snackbar.LENGTH_LONG).show();
//                                            }
//                                        });
//                                    } else {
//                                        Snackbar.make(view, "无商品可添加", Snackbar.LENGTH_LONG).show();
//                                    }
//                                }
//                            });
//                        } else if (menuItem.getItemId() == R.id.more_collect) {
//
//                        }
//                    }
//
//                    @Override
//                    public void onSheetDismissed(@NonNull BottomSheetMenuDialogFragment bottomSheetMenuDialogFragment, @Nullable Object o, int i) {
//
//                    }
//                }).show(getActivity().getSupportFragmentManager());
//            }
//        });
        binding.imageButtonSort.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                new BottomSheetMenuDialogFragment.Builder(getActivity()).setSheet(R.menu.menu_sort).setTitle("Sort").setListener(new BottomSheetListener() {
                    @Override
                    public void onSheetShown(@NonNull BottomSheetMenuDialogFragment bottomSheetMenuDialogFragment, @Nullable Object o) {

                    }

                    @Override
                    public void onSheetItemSelected(@NonNull BottomSheetMenuDialogFragment bottomSheetMenuDialogFragment, @NonNull MenuItem menuItem, @Nullable Object o) {
                        if (menuItem.getItemId() == R.id.sort_max_min) {
                            //从高到低
                            list.sort(new Comparator<ProductSimple>() {
                                @Override
                                public int compare(ProductSimple o1, ProductSimple o2) {
                                    if (Float.valueOf(o1.getPrice()) - Float.valueOf(o2.getPrice()) < 0) {
                                        return 1;
                                    }
                                    return -1;
                                }
                            });
                            adapter.notifyDataSetChanged();
                        } else if (menuItem.getItemId() == R.id.sort_min_max) {
                            //从低到高
                            list.sort(new Comparator<ProductSimple>() {
                                @Override
                                public int compare(ProductSimple o1, ProductSimple o2) {
                                    if (Float.valueOf(o1.getPrice()) - Float.valueOf(o2.getPrice()) > 0) {
                                        return 1;
                                    }
                                    return -1;
                                }
                            });
                            adapter.notifyDataSetChanged();
                        }
                    }

                    @Override
                    public void onSheetDismissed(@NonNull BottomSheetMenuDialogFragment bottomSheetMenuDialogFragment, @Nullable Object o, int i) {

                    }
                }).show(getActivity().getSupportFragmentManager());
            }
        });
        binding.imageButtonFilter.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                //直接新建一个碎片
            }
        });
        binding.sellPagerFramelayout.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                ARouter.getInstance().build("/mine_pager/app_search_goodsActivity").navigation();
                getActivity().overridePendingTransition(com.example.common.R.anim.set_in, com.example.common.R.anim.set_out);
            }
        });
        return view;
    }
}