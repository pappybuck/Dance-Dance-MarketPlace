using Microsoft.AspNetCore.Identity;
namespace Janus.Identity
{
    public class User : IdentityUser
    {
        public string FirstName { get; set; }
        public string LastName { get; set; }
        public List<RefreshToken?> token { get; set; }
        
        public Profile Profile { get; set; }

    }
}
