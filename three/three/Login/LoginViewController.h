//
//  LoginViewController.h
//  three
//
//  Created by Huc on 2023/12/11.
//

#import <UIKit/UIKit.h>

NS_ASSUME_NONNULL_BEGIN

@interface LoginViewController : UIViewController
<UITextFieldDelegate>

@property UITextField *mainTextField;
@property UITextField *subTextField;
@property UIButton *codeBtn;
@property UITextField *secureTextField;
@property UITextField *resecureTextField;
@property UIButton *loginBtn;
@property UIButton *signBtn;
@property UILabel *warnningLabel;
@property UIView* fakeView;
@property NSString *msg;
@property NSInteger countdownNum;

@end

NS_ASSUME_NONNULL_END
