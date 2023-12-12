//
//  LoginViewController.m
//  three
//
//  Created by Huc on 2023/12/11.
//

#import "LoginViewController.h"
#import <AudioToolbox/AudioToolbox.h>

@interface LoginViewController ()

@end

@implementation LoginViewController

- (void)viewDidLoad {
    [super viewDidLoad];
    // Do any additional setup after loading the view.
    [self loadBasicView];
    [self codeLoginView];
}

- (void)viewWillAppear:(BOOL)animated {
    self.tabBarController.tabBar.alpha = 1;
}

- (void)viewDidAppear:(BOOL)animated {
    self.tabBarController.tabBar.alpha = 0;
}

- (void)codeLoginView {
    self.mainTextField.placeholder = @"输入手机号";
    self.loginBtn.selected = NO;
    self.warnningLabel.text = @"";
    self.secureTextField.hidden = YES;
    self.resecureTextField.hidden = YES;
    self.codeBtn.hidden = NO;
    self.subTextField.hidden = NO;
    self.mainTextField.text = @"";
    self.subTextField.text = @"";
    self.secureTextField.text = @"";
    self.resecureTextField.text = @"";
    self.subTextField.frame = CGRectMake(self.view.bounds.size.width/2 - 150, self.view.bounds.size.height/2 + 120, 160, 40);
    self.codeBtn.frame = CGRectMake(self.view.bounds.size.width/2 + 20, self.view.bounds.size.height/2 + 120, 130, 40);
}

- (void)passwordLoginView {
    self.mainTextField.placeholder = @"输入账号";
    self.loginBtn.selected = YES;
    self.warnningLabel.text = @"";
    self.secureTextField.hidden = NO;
    self.resecureTextField.hidden = YES;
    self.codeBtn.hidden = YES;
    self.subTextField.hidden = YES;
    self.mainTextField.text = @"";
    self.subTextField.text = @"";
    self.secureTextField.text = @"";
    self.resecureTextField.text = @"";
    self.secureTextField.frame = CGRectMake(self.view.bounds.size.width/2 - 150, self.view.bounds.size.height/2 + 120, 300, 40);
}

- (void)registerView {
    self.mainTextField.placeholder = @"输入手机号";
    self.warnningLabel.text = @"";
    self.secureTextField.hidden = NO;
    self.resecureTextField.hidden = NO;
    self.codeBtn.hidden = NO;
    self.subTextField.hidden = NO;
    self.mainTextField.text = @"";
    self.subTextField.text = @"";
    self.secureTextField.text = @"";
    self.resecureTextField.text = @"";
    self.mainTextField.frame = CGRectMake(self.view.bounds.size.width/2 - 150, self.view.bounds.size.height/2 + 60, 300, 40);
    self.subTextField.frame = CGRectMake(self.view.bounds.size.width/2 - 150, self.view.bounds.size.height/2 + 240, 160, 40);
    self.codeBtn.frame = CGRectMake(self.view.bounds.size.width/2 + 20, self.view.bounds.size.height/2 + 240, 130, 40);
    self.secureTextField.frame = CGRectMake(self.view.bounds.size.width/2 - 150, self.view.bounds.size.height/2 + 120, 300, 40);
    self.resecureTextField.frame = CGRectMake(self.view.bounds.size.width/2 - 150, self.view.bounds.size.height/2 + 180, 300, 40);
}

- (void)loadBasicView {
    self.view.backgroundColor = UIColor.systemGray6Color;
//    UIImage* image = [UIImage imageNamed:@"淘牛马.jpeg"];
//    UIGraphicsBeginImageContext(self.view.bounds.size);
//    [image drawInRect:CGRectMake(0.0f, 0.0f, self.view.bounds.size.width, self.view.bounds.size.height)];
//    image = UIGraphicsGetImageFromCurrentImageContext();
//    UIGraphicsEndImageContext();
//    [self.view addSubview: [[UIImageView alloc] initWithImage:image]];
    self.mainTextField = [[UITextField alloc] init];
    self.subTextField = [[UITextField alloc] init];
    self.secureTextField = [[UITextField alloc] init];
    self.resecureTextField = [[UITextField alloc] init];
    self.codeBtn = [UIButton buttonWithType:UIButtonTypeCustom];
    self.signBtn = [UIButton buttonWithType:UIButtonTypeCustom];
    self.mainTextField.backgroundColor = UIColor.whiteColor;
    self.subTextField.backgroundColor = UIColor.whiteColor;
    self.secureTextField.backgroundColor = UIColor.whiteColor;
    self.resecureTextField.backgroundColor = UIColor.whiteColor;
    self.codeBtn.backgroundColor = UIColor.systemBlueColor;
    [self.codeBtn setTintColor:UIColor.whiteColor];
    self.signBtn.backgroundColor = UIColor.systemBlueColor;
    [self.signBtn setTintColor:UIColor.whiteColor];
    self.warnningLabel = [[UILabel alloc] init];
    
    self.fakeView = [[UIView alloc] init];
    
    self.loginBtn = [UIButton buttonWithType:UIButtonTypeCustom];
    UIButton *findBackBtn = [UIButton buttonWithType:UIButtonTypeCustom];
    NSMutableAttributedString *title = [[NSMutableAttributedString alloc] initWithString:@"密码登录"];
    NSRange titleRange = {0,[title length]};
    [title addAttribute:NSUnderlineStyleAttributeName value:[NSNumber numberWithInteger:NSUnderlineStyleSingle] range:titleRange];
    [self.loginBtn setAttributedTitle:title
                        forState:UIControlStateNormal];
//    [codeLoginBtn setTitle:@"验证登录" forState:UIControlStateNormal];
    title = [[NSMutableAttributedString alloc] initWithString:@"验证登录"];
    [title addAttribute:NSUnderlineStyleAttributeName value:[NSNumber numberWithInteger:NSUnderlineStyleSingle] range:titleRange];
    [self.loginBtn setAttributedTitle:title
                        forState:UIControlStateSelected];
    title = [[NSMutableAttributedString alloc] initWithString:@"找回密码"];
    [title addAttribute:NSUnderlineStyleAttributeName value:[NSNumber numberWithInteger:NSUnderlineStyleSingle] range:titleRange];
    [findBackBtn setAttributedTitle:title
                        forState:UIControlStateNormal];
//    [passwordLoginBtn setTitle:@"密码登录" forState:UIControlStateNormal];
    [self.loginBtn setTitleColor:UIColor.systemBlueColor forState:UIControlStateNormal];
    [findBackBtn setTitleColor:UIColor.systemBlueColor forState:UIControlStateNormal];
    [self.loginBtn addTarget:self action:@selector(pressLoginBtn:) forControlEvents:UIControlEventTouchUpInside];
    [findBackBtn addTarget:self action:@selector(pressFindBackBtn) forControlEvents:UIControlEventTouchUpInside];
    self.loginBtn.frame = CGRectMake(self.view.bounds.size.width/2, self.view.bounds.size.height/2 + 180, 150, 40);
    findBackBtn.frame = CGRectMake(self.view.bounds.size.width/2 - 150, self.view.bounds.size.height/2 + 180, 150, 40);
    [self.view addSubview:self.loginBtn];
    [self.view addSubview:findBackBtn];
    
    self.mainTextField.frame = CGRectMake(self.view.bounds.size.width/2 - 150, self.view.bounds.size.height/2 + 60, 300, 40);
    self.subTextField.frame = CGRectMake(self.view.bounds.size.width/2 - 150, self.view.bounds.size.height/2 + 240, 160, 40);
    self.codeBtn.frame = CGRectMake(self.view.bounds.size.width/2 + 20, self.view.bounds.size.height/2 + 240, 130, 40);
    self.secureTextField.frame = CGRectMake(self.view.bounds.size.width/2 - 150, self.view.bounds.size.height/2 + 120, 300, 40);
    self.resecureTextField.frame = CGRectMake(self.view.bounds.size.width/2 - 150, self.view.bounds.size.height/2 + 180, 300, 40);
    self.signBtn.frame = CGRectMake(self.view.bounds.size.width/2 - 150, self.view.bounds.size.height/2 + 300, 300, 40);
    self.warnningLabel.frame = CGRectMake(self.view.bounds.size.width/2 - 150, self.view.bounds.size.height/2 + 30, 300, 20);
    
    [self.codeBtn addTarget:self action:@selector(pressCodeBtn) forControlEvents:UIControlEventTouchUpInside];
    [self.signBtn addTarget:self action:@selector(pressSignBtn:) forControlEvents:UIControlEventTouchUpInside];
    
    self.mainTextField.placeholder = @"输入手机号";
    self.subTextField.placeholder = @"输入验证码";
    self.secureTextField.placeholder = @"输入密码";
    self.resecureTextField.placeholder = @"重复密码";
    self.secureTextField.secureTextEntry = YES;
    self.resecureTextField.secureTextEntry = YES;
    [self.codeBtn setTitle:@"发送验证码" forState:UIControlStateNormal];
    [self.signBtn setTitle:@"登录" forState:UIControlStateNormal];
    self.warnningLabel.textColor = UIColor.redColor;
    self.warnningLabel.textAlignment = NSTextAlignmentLeft;
    self.mainTextField.delegate = self;
    self.subTextField.delegate = self;
    [self.view addSubview:self.mainTextField];
    [self.view addSubview:self.subTextField];
    [self.view addSubview:self.secureTextField];
    [self.view addSubview:self.resecureTextField];
    [self.view addSubview:self.codeBtn];
    [self.view addSubview:self.signBtn];
    [self.view addSubview:self.warnningLabel];
    
    UISegmentedControl *segmentedControl = [[UISegmentedControl alloc] init];
    segmentedControl.frame = CGRectMake(self.view.bounds.size.width/2 - 140, self.view.bounds.size.height/2 - 20, 160, 40);
    [segmentedControl insertSegmentWithTitle: @"登录" atIndex: 0 animated: NO];
    [segmentedControl insertSegmentWithTitle: @"注册" atIndex: 1 animated: NO];
    segmentedControl.selectedSegmentIndex = 0;
    [segmentedControl addTarget: self action: @selector(segmentedControlChange:) forControlEvents: UIControlEventValueChanged];
    segmentedControl.backgroundColor = UIColor.systemGray6Color;
    [self.view addSubview: segmentedControl];
    
    [[NSNotificationCenter defaultCenter]addObserver:self selector:@selector(keyboardWillShow:) name:UIKeyboardWillShowNotification object:nil];
    [[NSNotificationCenter defaultCenter]addObserver:self selector:@selector(keyboardWillHide:) name:UIKeyboardWillHideNotification object:nil];
}

- (void)pressLoginBtn:(UIButton *)btn {
    btn.selected = !btn.selected;
    if (btn.selected) {
        [self passwordLoginView];
    } else {
        [self codeLoginView];
    }
}

- (void)pressFindBackBtn {
    
}

-(void)keyboardWillShow:(NSNotification*)note{
    CGFloat boundX = self.view.bounds.size.width;
    CGFloat boundY = self.view.bounds.size.height;
    self.view.frame = CGRectMake(0, -boundY/2 + 160, boundX, boundY);

}

-(void)keyboardWillHide:(NSNotification*)note{
    CGFloat boundX = self.view.bounds.size.width;
    CGFloat boundY = self.view.bounds.size.height;
    self.view.frame = CGRectMake(0, 0, boundX, boundY);
}

- (void)sendMessage:(NSDictionary *)message withPurpose:(NSString *)purpose {
    NSString *str = [NSString stringWithFormat:@"http://192.168.0.146:13000/%@", purpose];
    NSURL *url = [NSURL URLWithString:str];
    NSMutableURLRequest *request = [NSMutableURLRequest requestWithURL:url];
    request.HTTPMethod = @"POST";

    // 创建要发送的JSON参数
    NSError *error;
    NSData *jsonData = [NSJSONSerialization dataWithJSONObject:message options:0 error:&error];
    if (error) {
        NSLog(@"Failed to serialize JSON: %@", error);
        return;
    }

    // 设置请求体为JSON参数
    request.HTTPBody = jsonData;
    [request setValue:@"application/json" forHTTPHeaderField:@"Content-Type"];

    NSURLSession *session = [NSURLSession sharedSession];
    NSURLSessionDataTask *dataTask = [session dataTaskWithRequest:request completionHandler:^(NSData *data, NSURLResponse *response, NSError *error) {
        if (error) {
            NSLog(@"Failed to send request: %@", error);
            return;
        }
        // 处理响应数据
        NSString *responseData = [[NSString alloc] initWithData:data encoding:NSUTF8StringEncoding];
        NSLog(@"Response: %@", responseData);
        NSDictionary *dic = [NSJSONSerialization JSONObjectWithData:
                                 data options:kNilOptions error:nil];
        if ([dic[@"code"] intValue] > 201) {
            self.msg = [NSString stringWithFormat:@"！%@", dic[@"msg"]];
        } else {
            dispatch_async(dispatch_get_main_queue(), ^{
                if ([dic[@"msg"] isEqualToString:@"登录成功"]) {
                    UIAlertController* alertView = [UIAlertController alertControllerWithTitle:dic[@"msg"]
                    message:nil
                    preferredStyle:UIAlertControllerStyleAlert];

                    
                    UIAlertAction* cancel = [UIAlertAction actionWithTitle:@"确定"
                    style:UIAlertActionStyleCancel handler:nil];
                    
                    [alertView addAction: cancel];
                         

                    // 显示视图
                    [self presentViewController: alertView animated:YES
                            completion:nil];
                    self.navigationController.tabBarController.title = dic[@"data"][@"token"];
                    NSLog(@"%@", self.navigationController.tabBarController.title);
                } else if ([dic[@"msg"] isEqualToString:@"登录成功"]) {
                    UIAlertController* alertView = [UIAlertController alertControllerWithTitle:dic[@"msg"]
                    message:dic[@"data"]
                    preferredStyle:UIAlertControllerStyleAlert];

                    
                    UIAlertAction* cancel = [UIAlertAction actionWithTitle:@"确定"
                    style:UIAlertActionStyleCancel handler:nil];
                    
                    [alertView addAction: cancel];
                         

                    // 显示视图
                    [self presentViewController: alertView animated:YES
                            completion:nil];
                }
            });
        }
        [self changeWarnLabel];
    }];
    [dataTask resume];
}

- (void)changeWarnLabel {
    dispatch_async(dispatch_get_main_queue(), ^{
        self.warnningLabel.text = self.msg;
        self.msg = @"";
    });
}

- (void)pressCodeBtn {
    if (self.codeBtn.selected) {
        return;
    }
    if (self.mainTextField.text.length) {
        if (self.mainTextField.text.length != 11) {
            AudioServicesPlaySystemSound(kSystemSoundID_Vibrate);
            self.warnningLabel.text = @"！手机号输入错误";
        } else {
            NSDictionary *message = @{@"phone_number": self.mainTextField.text};
            [self sendMessage:message withPurpose:@"send_phone_code"];
            self.codeBtn.selected = YES;
            self.countdownNum = 60;
            NSTimer *sendCodeTimer = [NSTimer scheduledTimerWithTimeInterval:1 target:self selector:@selector(countDownCodeBtn:) userInfo:nil repeats:YES];
            NSString* str = [NSString stringWithFormat:@"重新发送（%lds）", self.countdownNum];
            [self.codeBtn setTitle:str forState:UIControlStateSelected];
            self.codeBtn.backgroundColor = UIColor.systemGrayColor;
        }
    } else {
        AudioServicesPlaySystemSound(kSystemSoundID_Vibrate);
        self.warnningLabel.text = @"！手机号不能为空";
    }
}

- (void)countDownCodeBtn:(NSTimer *)timer {
    self.countdownNum--;
    if (!self.countdownNum) {
        self.codeBtn.userInteractionEnabled = YES;
        self.codeBtn.backgroundColor = UIColor.systemBlueColor;
        self.codeBtn.selected = NO;
        [timer invalidate];
    } else {
        self.codeBtn.userInteractionEnabled = NO;
        NSString* str = [NSString stringWithFormat:@"重新发送（%lds）", self.countdownNum];
        [self.codeBtn setTitle:str forState:UIControlStateSelected];
        self.codeBtn.backgroundColor = UIColor.systemGrayColor;
    }
    
}

- (void)pressSignBtn:(UIButton *)btn {
    self.warnningLabel.text = @"";
    if ([btn.titleLabel.text isEqualToString:@"登录"]) {
        if (self.mainTextField.text.length) {
            if (self.loginBtn.selected) {
                if (self.mainTextField.text.length != 11) {
                    AudioServicesPlaySystemSound(kSystemSoundID_Vibrate);
                    self.warnningLabel.text = @"！账号输入错误";
                } else if (self.secureTextField.text.length) {
                    NSDictionary *message = @{@"account":self.mainTextField.text, @"password":self.secureTextField.text};
                    [self sendMessage:message withPurpose:@"user/login/password"];
                } else {
                    AudioServicesPlaySystemSound(kSystemSoundID_Vibrate);
                    self.warnningLabel.text = @"！密码不能为空";
                }
            } else {
                if (self.mainTextField.text.length != 11) {
                    AudioServicesPlaySystemSound(kSystemSoundID_Vibrate);
                    self.warnningLabel.text = @"！手机号输入错误";
                } else if (self.subTextField.text.length) {
                    NSDictionary *message = @{@"phone_number":self.mainTextField.text, @"verification_code":self.subTextField.text};
                    [self sendMessage:message withPurpose:@"user/login/phone"];
                } else {
                    AudioServicesPlaySystemSound(kSystemSoundID_Vibrate);
                    self.warnningLabel.text = @"！验证码不能为空";
                }
            }
        } else {
            AudioServicesPlaySystemSound(kSystemSoundID_Vibrate);
            if (self.loginBtn.selected) {
                self.warnningLabel.text = @"！账号不能为空";
            } else {
                self.warnningLabel.text = @"！手机号不能为空";
            }
        }
    } else {
        if (self.mainTextField.text.length) {
            if (self.subTextField.text.length) {
                if (self.secureTextField.text.length) {
                    if (self.resecureTextField.text.length) {
                        if ([self.resecureTextField.text isEqualToString:self.secureTextField.text]) {
                            NSDictionary *message = @{@"phone_number":self.mainTextField.text, @"password":self.secureTextField.text, @"verification_code":self.subTextField.text};
                            [self sendMessage:message withPurpose:@"user/register/phone"];
                        } else {
                            AudioServicesPlaySystemSound(kSystemSoundID_Vibrate);
                            self.warnningLabel.text = @"！两次密码不一样";
                        }
                    } else {
                        AudioServicesPlaySystemSound(kSystemSoundID_Vibrate);
                        self.warnningLabel.text = @"！请重复输入密码";
                    }
                } else {
                    AudioServicesPlaySystemSound(kSystemSoundID_Vibrate);
                    self.warnningLabel.text = @"！密码不能为空";
                }
            } else {
                AudioServicesPlaySystemSound(kSystemSoundID_Vibrate);
                self.warnningLabel.text = @"！验证码不能为空";
            }
        } else {
            AudioServicesPlaySystemSound(kSystemSoundID_Vibrate);
            self.warnningLabel.text = @"！手机号不能为空";
        }
    }
}

- (void)segmentedControlChange:(UISegmentedControl *)segmentedControl {
    if (segmentedControl.selectedSegmentIndex) {
        [self registerView];
        [self.signBtn setTitle:@"注册" forState:UIControlStateNormal];
    } else {
        [self codeLoginView];
        [self.signBtn setTitle:@"登录" forState:UIControlStateNormal];
    }
}

- (BOOL)textField:(UITextField *)textField shouldChangeCharactersInRange:(NSRange)range replacementString:(NSString *)string {
    [self changeWarnLabel];
    if (textField == self.mainTextField || textField == self.subTextField) {
        return [self validateNumber: string];
    }
    return YES;
}

- (BOOL)validateNumber:(NSString*)number {
    BOOL res = YES;
    NSCharacterSet* tmpSet = [NSCharacterSet characterSetWithCharactersInString:@"0123456789"];
    int i = 0;
    while (i < number.length) {
        NSString * string = [number substringWithRange:NSMakeRange(i, 1)];
        NSRange range = [string rangeOfCharacterFromSet:tmpSet];
        if (range.length == 0) {
            res = NO;
            break;
        }
        i++;
    }
    return res;
}

- (void)searchBarTextDidBeginEditing:(UISearchBar *)searchBar{
    self.fakeView.backgroundColor = UIColor.clearColor;
    self.fakeView.frame = self.view.frame;
    [self.view addSubview: self.fakeView];
}

- (void) touchesBegan:(NSSet<UITouch *> *)touches withEvent:(UIEvent *)event {
    [self.mainTextField resignFirstResponder];
    [self.subTextField resignFirstResponder];
    [self.secureTextField resignFirstResponder];
    [self.resecureTextField resignFirstResponder];
    [self.fakeView removeFromSuperview];
}

/*
#pragma mark - Navigation

// In a storyboard-based application, you will often want to do a little preparation before navigation
- (void)prepareForSegue:(UIStoryboardSegue *)segue sender:(id)sender {
    // Get the new view controller using [segue destinationViewController].
    // Pass the selected object to the new view controller.
}
*/

@end
