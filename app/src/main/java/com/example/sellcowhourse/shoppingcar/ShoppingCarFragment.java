package com.example.sellcowhourse.shoppingcar;

import android.annotation.SuppressLint;
import android.content.Context;
import android.os.Bundle;
import android.os.Handler;
import android.os.HandlerThread;
import android.util.DisplayMetrics;
import android.view.LayoutInflater;
import android.view.MotionEvent;
import android.view.View;
import android.view.ViewGroup;
import android.view.WindowManager;
import android.view.animation.Animation;
import android.view.animation.ScaleAnimation;
import android.widget.Toast;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.recyclerview.widget.GridLayoutManager;
import androidx.swiperefreshlayout.widget.SwipeRefreshLayout;

import com.example.common.room.AppDatabase;
import com.example.common.room.dao.ShoppingCarDao;
import com.example.common.room.entitues.ShoppingCarOrder;
import com.example.sellcowhourse.R;
import com.example.sellcowhourse.databinding.FragmentShoppingCarBinding;
import com.example.sellcowhourse.shoppingcar.adapter.ShoppingCarOrderAdapter;

import java.util.ArrayList;
import java.util.List;
import java.util.function.Consumer;

/**
 * @Author winiymissl
 * @Date 2024-02-20 20:45
 * @Version 1.0
 */
public class ShoppingCarFragment extends Fragment {
    FragmentShoppingCarBinding binding;
    private int screenWidth = 720;
    private int screenHeight = 1280;
    List<ShoppingCarOrder> list = new ArrayList<>();
    @SuppressLint("ClickableViewAccessibility")
    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_shopping_car, container, false);
        binding = FragmentShoppingCarBinding.bind(view);
        ShoppingCarOrderAdapter adapter = new ShoppingCarOrderAdapter(list, getActivity());
        GridLayoutManager gridLayoutManager = new GridLayoutManager(getActivity(), 1);
        binding.LrecyclerViewShoppingCar.setLayoutManager(gridLayoutManager);
        binding.LrecyclerViewShoppingCar.setAdapter(adapter);
        HandlerThread handlerThread = new HandlerThread("Refresh");
        handlerThread.start();
        Handler handler = new Handler(handlerThread.getLooper());
        handler.post(new Runnable() {
            @Override
            public void run() {
                ShoppingCarDao shoppingCarDao = AppDatabase.getInstance(getActivity()).shoppingCarDao();
                shoppingCarDao.getAllInfo().forEach(new Consumer<ShoppingCarOrder>() {
                    @Override
                    public void accept(ShoppingCarOrder shoppingCarOrder) {
                        list.add(shoppingCarOrder);
                    }
                });

                getActivity().runOnUiThread(new Runnable() {
                    @Override
                    public void run() {
                        adapter.notifyDataSetChanged();
                    }
                });
            }
        });

        binding.swipeRefreshLayoutShoppingCar.setOnRefreshListener(new SwipeRefreshLayout.OnRefreshListener() {
            @Override
            public void onRefresh() {
                synchronized (new Object()) {
                    HandlerThread handlerThread = new HandlerThread("Refresh");
                    handlerThread.start();
                    Handler handler = new Handler(handlerThread.getLooper());
                    handler.post(new Runnable() {
                        @Override
                        public void run() {
                            list.clear();
                            ShoppingCarDao shoppingCarDao = AppDatabase.getInstance(getActivity()).shoppingCarDao();
                            List<ShoppingCarOrder> temp = shoppingCarDao.getAllInfo();
                            list.addAll(temp);
                            getActivity().runOnUiThread(new Runnable() {
                                @Override
                                public void run() {
                                    adapter.notifyDataSetChanged();
                                    binding.swipeRefreshLayoutShoppingCar.setRefreshing(false);
                                }
                            });
                        }
                    });
                }
            }
        });
        binding.floatingButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Animation anim = new ScaleAnimation(1, 0.8f, 1, 0.8f, Animation.RELATIVE_TO_SELF, 0.5f, Animation.RELATIVE_TO_SELF, 0.5f);
                anim.setDuration(100); // 设置动画持续时间
                anim.setRepeatMode(Animation.REVERSE); // 设置动画重复模式
                anim.setRepeatCount(1); // 设置动画重复次数
                // 应用动画到 FloatingActionButton
                binding.floatingButton.startAnimation(anim);
                //结算，生成订单
                Toast.makeText(getActivity(), "点击事件", Toast.LENGTH_SHORT).show();
                //进行碎片操作
            }
        });
        WindowManager wm = (WindowManager) getActivity().getSystemService(Context.WINDOW_SERVICE);
        DisplayMetrics dm = new DisplayMetrics();
        wm.getDefaultDisplay().getMetrics(dm);
        screenWidth = dm.widthPixels;
        screenHeight = dm.heightPixels;
        binding.floatingButton.setOnTouchListener(new View.OnTouchListener() {
            private float lastX = 0;
            private float lastY = 0;
            private float beginX = 0;
            private float beginY = 0;

            @Override
            public boolean onTouch(View v, MotionEvent event) {

                switch (event.getAction()) {
                    case MotionEvent.ACTION_DOWN:
                        lastX = (int) event.getRawX();   // 触摸点与屏幕左边的距离
                        lastY = (int) event.getRawY();   // 触摸点与屏幕上边的距离
                        beginX = lastX;
                        beginY = lastY;
                        break;
                    case MotionEvent.ACTION_MOVE:

                        float dx = event.getRawX() - lastX;    // x轴拖动的绝对距离
                        float dy = event.getRawY() - lastY;    // y轴拖动的绝对距离

                        // getLeft(): 子View的左边界到父View的左边界的距离, getRight():子View的右边界到父View的左边界的距离
                        // 如下几个数据表示view应该在布局中的位置
                        float left = v.getLeft() + dx;
                        float top = v.getTop() + dy;
                        float right = v.getRight() + dx;
                        float bottom = v.getBottom() + dy;
                        if (left < 0) {
                            left = 0;
                            right = left + v.getWidth();
                        }
                        if (right > screenWidth) {
                            right = screenWidth;
                            left = right - v.getWidth();
                        }

                        if (top < 0) {
                            top = 0;
                            bottom = top + v.getHeight();
                        }
                        if (bottom > screenHeight) {
                            bottom = screenHeight;
                            top = bottom - v.getHeight();
                        }
                        //修改现在的位置，没有改变布局参数，所以通过requestLayout()方法，
                        // 会回到原位，每次返回碎片的时候，就会回到原位
                        v.layout((int) left, (int) top, (int) right, (int) bottom);
                        lastX = event.getRawX();
                        lastY = event.getRawY();
                        break;
                    case MotionEvent.ACTION_UP:

                        // 解决拖拽的时候松手点击事件触发
                        if (Math.abs(lastX - beginX) < 10 && Math.abs(lastY - beginY) < 10) {
                            return v.onTouchEvent(event);//等同于返回false
                        } else {
                            v.setPressed(false);
                            return true;
                        }

                    default:
                        break;
                }
                return false;
            }

        });
        return view;
    }
}

